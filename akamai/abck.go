package akamai

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func randomNumberInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func randomNumberFloat() float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Float64()
}

func FpVal(bmak *Bmak) Fpcf {
	return Fpcf{
		RVal: fmt.Sprintf("%.f", math.Floor(randomNumberFloat()*(1000-1+1))+1),
		RCfp: "-1249560680",
		Td:   math.Floor(randomNumberFloat()*(23-17+1)) + 17,
	}

}

func DeviceAct(bmak *Bmak) *Bmak {
	do := math.Floor(randomNumberFloat()*(50-15+1)) + 15

	if do < 25 {
		bmak.Doact = ""
		bmak.DoeVel = 0
		bmak.Dmact = ""
		bmak.DmeVel = 0
		bmak.Ta = 0
	} else {
		DoActTime := math.Floor(randomNumberFloat()*(310-50+1)) + 50
		bmak.Doact = "0," + fmt.Sprintf("%.f", DoActTime) + ",-1,-1,-1;"
		bmak.DoeVel = DoActTime
		bmak.Ta = bmak.Ta + DoActTime

		DMacTime := DoActTime - math.Floor(randomNumberFloat()*(60-5+1)) + 5
		bmak.Dmact = "0," + fmt.Sprintf("%.f", DMacTime) + ",-1,-1,-1,-1,-1,-1,-1,-1,-1;"
		bmak.DmeVel = DMacTime
		bmak.Ta = bmak.Ta + DMacTime
	}

	return bmak
}

func GetCfDate() int64 {
	t := time.Now()
	tUnixMilli := int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond)
	return tUnixMilli
}

func D3() int64 {
	return GetCfDate() % 1e7
}

func rir(t int, a int, e int, n int) int {
	if t > a && t <= e {
		t += n % (e - a)
		if t > e {
			t = t - e + a
		}
	}
	return t
}

func Od(t string, a string) string {
	var e []string
	n := len(a)
	if n > 0 {
		for o := 0; o < len(t); o++ {
			m := []rune(t)[o]
			r := string(t[o])
			i := a[(o % n)]
			m = rune(rir(int(m), 47, 57, int(i)))
			if m != []rune(t)[o] {
				r = string(m)
			}
			e = append(e, r)
		}
		if len(e) > 0 {
			return strings.Join(e[:], "")
		}
	}
	return t
}

func Ab(t string) int {
	a := 0
	for e := 0; e < len(t); e++ {
		n := []rune(t)[e]
		if n < 128 {
			a += int(n)
		}
	}
	return a
}

func Uar(userAgent string) string {
	return regexp.MustCompile(`/\\|"/g`).ReplaceAllString(userAgent, "")
}

func Fas(browser string) (int, error) {
	if browser == "chrome" {
		return 30261693, nil
	} else if browser == "firefox" {
		return 26067385, nil
	} else {
		return 0, errors.New("Unsupported browser specified (use \"chrome\" or \"firefox\")")
	}
}

func mrjp(flt float64) float64 {
	return flt
}

func mrinf(flt float64) bool {
	return math.IsInf(flt, 0)
}

func invoke(fn interface{}, args ...float64) {
	v := reflect.ValueOf(fn)
	rargs := make([]reflect.Value, len(args))
	for i, a := range args {
		rargs[i] = reflect.ValueOf(a)
	}
	v.Call(rargs)
}

func Getmr() string {
	t := ""
	var e []interface{}
	e = append(e, math.Abs, math.Acos, math.Asin, math.Atanh, math.Cbrt, math.Exp, mrjp, math.Sqrt, mrinf, math.IsNaN, mrjp, mrjp, mrjp, mrjp)
	for n := 0; n < len(e); n++ {
		var o []float64
		var m float64
		m = 0.0
		r := time.Now()
		c := 0
		for i := 0; i < 1000 && m < .6; i++ {
			b := time.Now()
			for d := 0; d < 4e3; d++ {
				invoke(e[n], 3.14)
			}
			k := time.Now()
			o = append(o, 1000*math.Round(float64(k.UnixNano()/1000000)-float64(b.UnixNano()/1000000)))
			m = float64(k.UnixNano()/1000000) - float64(r.UnixNano()/1000000)
			s := o
			sort.Float64s(s)
			c = int(s[int(math.Floor(float64(len(s)/2)))]) / 5
			t = fmt.Sprintf("%v%v,", t, c)
		}
	}
	return t
}

func Sed(browser string) string {
	return "0,0,0,0,1,0,0"
}

func Np(browser string) (string, error) {
	if browser == "chrome" {
		return "10321144241322243122", nil
	} else if browser == "firefox" {
		return "11133333331333333333", nil
	} else {
		return "", errors.New("Unsupported browser specified (use \"chrome\" or \"firefox\")")
	}
}

func calDis(t []int) float64 {
	a := t[0] - t[1]
	e := t[2] - t[3]
	n := t[4] - t[5]
	return math.Floor(math.Sqrt(float64(a*a + e*e + n*n)))
}

func Jrs(t int) ([]float64, error) {
	a := int(math.Floor(100000*randomNumberFloat() + 10000))
	e := strconv.Itoa(t * a)
	m := len(e) >= 18
	n := 0
	var o []int
	for len(o) < 6 {
		num, err := strconv.Atoi(e[n:(n + 2)])
		if err != nil {
			return []float64{}, err
		}
		o = append(o, num)
		if m {
			n = n + 3
		} else {
			n = n + 2
		}
	}
	return []float64{float64(a), calDis(o)}, nil
}

func cc(t int) func(int, int) int {
	a := t % 4
	if 2 == a {
		a = 3
	}
	e := 42 + a
	var n func(int, int) int
	n = func(t int, a int) int {
		return 0
	}
	if 42 == e {
		n = func(t int, a int) int {
			return t * a
		}
	} else if 43 == e {
		n = func(t int, a int) int {
			return t + a
		}
	} else {
		n = func(t int, a int) int {
			return t - a
		}
	}
	return n
}

func O9(d3 int64) int {
	a := int(d3)
	for n := 0; n < 5; n++ {
		o := math.Mod(math.Floor(float64(d3)/math.Pow(10, float64(n))), 10)
		m := o + 1
		op := cc(int(o))
		a = op(int(a), int(m))
	}

	return a * 3
}

func GetType(t string) int {
	t = strings.ToLower(t)
	if "text" == t || "search" == t || "url" == t || "email" == t || "tel" == t || "number" == t {
		return 0
	} else if "password" == t {
		return 1
	} else {
		return 2
	}
}

func Z1(startTimestamp int64) int64 {
	return startTimestamp / 2016 * 2016
}

func floatToString(flt float64, precision int) string {
	return strconv.FormatFloat(flt, 'f', precision, 64)
}

type ScreenData struct {
	AvailWidth  int
	AvailHeight int
	Width       int
	Height      int
	InnerWidth  int
	InnerHeight int
	OuterWidth  int
}

type BrowserData struct {
	Lang       string
	UserAgent  string
	ScreenData ScreenData
}

func Gd(browser string, browserData BrowserData, startTimestamp int64, d3 int64) (string, error) {
	userAgent := Uar(browserData.UserAgent)
	d := randomNumberFloat()
	xagg := 0
	psub := 0
	plen := 0
	bd := ""
	if browser == "chrome" {
		xagg = 12147
		psub = 20030107
		plen = 3
		bd = ",cpen:0,i1:0,dm:0,cwen:0,non:1,opc:0,fc:0,sc:0,wrc:1,isc:0,vib:1,bat:1,x11:0,x12:1"
	} else if browser == "firefox" {
		xagg = 11059
		psub = 20100101
		plen = 0
		bd = fmt.Sprintf(",cpen:0,i1:0,dm:0,cwen:0,non:1,opc:0,fc:1,sc:0,wrc:1,isc:%v,vib:1,bat:0,x11:0,x12:1", randomNumberInt(60, 100))
	} else {
		return "", errors.New("Unsupported browser specified (use \"chrome\" or \"firefox\")")
	}
	return fmt.Sprintf(
		"%v,uaend,%v,%v,%v,Gecko,%v,0,0,0,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,loc:",
		userAgent,
		xagg,
		psub,
		browserData.Lang,
		plen,
		Z1(startTimestamp),
		d3,
		browserData.ScreenData.AvailWidth,
		browserData.ScreenData.AvailHeight,
		browserData.ScreenData.Width,
		browserData.ScreenData.Height,
		browserData.ScreenData.InnerWidth,
		browserData.ScreenData.InnerHeight,
		browserData.ScreenData.OuterWidth,
		bd,
		Ab(userAgent),
		floatToString(d, 16)[0:11]+floatToString(math.Floor(1000*d/2), -1),
		startTimestamp/2), nil
}

func GetH(bmak *Bmak) string {
	d2 := int(bmak.Z1 / 23)
	log.Print(bmak.Z1 / 23)
	f := int(d2 / 6)
	updateT := GetCfDate() - bmak.StartTimeStamp
	startDif := GetCfDate() - bmak.StartTimeStamp
	u, err := Jrs(int(bmak.StartTimeStamp))
	if err != nil {
		log.Fatal(err)
	}

	var velSum = bmak.KeVel + bmak.MeVel + bmak.DoeVel + bmak.DmeVel + bmak.TeVel + bmak.PeVel

	fas, err := Fas("chrome")
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v", bmak.KeVel+1, bmak.MeVel+32, bmak.TeVel+32, bmak.DoeVel, bmak.DmeVel, bmak.PeVel, velSum, updateT, bmak.InitTime, bmak.StartTimeStamp, bmak.Fpcf.Td, d2, bmak.KeCnt, bmak.MeCnt, f, bmak.PeCnt, bmak.TeCnt, startDif, bmak.Ta, bmak.NCk, bmak.Abck, Ab(bmak.Abck), bmak.Fpcf.RVal, bmak.Fpcf.RCfp, fas, "PiZtE", u[0], u[1])
}
