package ratelimiter

import (
	"errors"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karlseguin/ccache"
	"golang.org/x/time/rate"
)

const DefaultCacheTTL = time.Minute * 5

var (
	RateErr = errors.New("too many requests")
)

type store interface {
	Get(string) (interface{}, bool)
	Set(string, interface{})
}

type localCache struct {
	c *ccache.Cache
}

func newLocalCache() *localCache {
	return &localCache{
		c: ccache.New(ccache.Configure()),
	}
}

func (lc *localCache) Get(k string) (interface{}, bool) {
	item := lc.c.Get(k)
	if item == nil || item.Expired() {
		return nil, false
	}
	return item.Value(), true
}

func (lc *localCache) Set(k string, v interface{}) {
	lc.c.Set(k, v, DefaultCacheTTL)
}

type Limiter struct {
	sync.Mutex
	reqStore   store
	RateLimit  rate.Limit
	BucketSize int
}

func New(limit float64, size int, reqStore store) *Limiter {
	if reqStore == nil {
		reqStore = newLocalCache()
	}
	return &Limiter{
		Mutex:      sync.Mutex{},
		reqStore:   reqStore,
		RateLimit:  rate.Limit(limit),
		BucketSize: size,
	}
}

func (gl *Limiter) getReqLimiter(rid string) *rate.Limiter {
	l, ok := gl.reqStore.Get(rid)
	if ok {
		return l.(*rate.Limiter)
	}

	newL := rate.NewLimiter(gl.RateLimit, gl.BucketSize)
	gl.reqStore.Set(rid, newL)
	return newL
}

// Wrapper around the time/rate Allow func that first retrives the
// limiter for a given request id (ip address) from the cache.
func (gl *Limiter) isAllow(rid string) bool {
	gl.Lock()
	l := gl.getReqLimiter(rid)
	gl.Unlock()

	return l.Allow()
}

// RateLimit gin middleware that limit the requests per second per context.
// Each request is identified by the ip address and cached in the limiter,
// until dropped by the lru cache or until expired.
func (gl *Limiter) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipAddr := extractIPAddr(c)

		if !gl.isAllow(ipAddr) {
			c.AbortWithError(http.StatusTooManyRequests, RateErr)
			return
		}
		c.Next()
	}
}

func extractIPAddr(c *gin.Context) string {
	ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err != nil {
		// TODO: figure out a fallback in cases where ip parsing from remote addr fails.
		return ""
	}
	return ip
}
