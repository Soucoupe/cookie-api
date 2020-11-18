package main

import (
	"net/http"

	"github.com/GuapAIO/akamai-api/akamai"
	"github.com/GuapAIO/akamai-api/config"
	"github.com/GuapAIO/akamai-api/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

type GeneratorRequest struct {
	LicenseKey string `json:"licenseKey,omitempty"`

	Abck string `json:"abck"`
	URL  string `json:"url"`
}

func main() {
	config.LoadConfig()

	dbPostgres := db.ConnectDb(&pg.Options{
		Addr:     config.Conf.Database.Addr,
		User:     config.Conf.Database.Username,
		Password: config.Conf.Database.Password,
		Database: config.Conf.Database.Database,
	})

	r := gin.Default()

	// CORS
	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/gen/akamai", func(c *gin.Context) {
		var json GeneratorRequest
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		randomSensor, err := db.GetRandomSensor(db)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		sensorData, _, err := akamai.GenCookie(json.URL, json.Abck)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"sensor_data": sensorData})
	})

	r.POST("/collector", func(c *gin.Context) {
		var json db.SensorData
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := db.InsertSensor(json, dbPostgres)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "thanks you!"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
