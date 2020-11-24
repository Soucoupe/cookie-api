package main

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/GuapAIO/akamai-api/akamai"
	"github.com/GuapAIO/akamai-api/config"
	"github.com/GuapAIO/akamai-api/db"
	"github.com/GuapAIO/akamai-api/px"
	"github.com/GuapAIO/akamai-api/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

type GeneratorRequest struct {
	LicenseKey string `json:"licenseKey,omitempty"`

	Payload string `json:"payload"`
}

type AkamaiGeneratorRequest struct {
	Abck string `json:"abck"`
	URL  string `json:"url"`
}

type PXGeneratorRequest struct {
	TargetURL string `json:"target"`
	URL       string `json:"url"`
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

	// Enable CORS
	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/gen/akamai", func(c *gin.Context) {
		// Unmarshal the JSON
		var genJSON GeneratorRequest
		if err := c.ShouldBindJSON(&genJSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Decode the base64 string
		decodedPayload, err := base64.RawStdEncoding.DecodeString(genJSON.Payload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Decrypt the decoded base64 payload with our secret encryption key
		decryptedPayload, err := utils.Decrypt(string(decodedPayload), utils.EncryptionKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Unmarhsal the real gen request!
		var akamaiJSON AkamaiGeneratorRequest
		err = json.Unmarshal(decryptedPayload, &akamaiJSON)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Get a random sensor in the DB
		randomSensor, err := db.GetRandomSensor(dbPostgres)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gen our Akamai Cookie
		sensorData, err := akamai.GenCookie(akamaiJSON.URL, akamaiJSON.Abck, *randomSensor)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Encrypt our Payload back
		encryptedPayload, err := utils.Encrypt(sensorData, utils.EncryptionKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Encode our encrypted payload in base64
		sEnc := base64.RawStdEncoding.EncodeToString([]byte(encryptedPayload))

		c.JSON(http.StatusOK, gin.H{"sensor_data": sEnc})
	})

	r.POST("/gen/px", func(c *gin.Context) {
		var genJSON GeneratorRequest
		if err := c.ShouldBindJSON(&genJSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Decode the base64 string
		decodedPayload, err := base64.RawStdEncoding.DecodeString(genJSON.Payload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Decrypt the decoded base64 payload with our secret encryption key
		decryptedPayload, err := utils.Decrypt(string(decodedPayload), utils.EncryptionKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var pxJSON PXGeneratorRequest
		err = json.Unmarshal(decryptedPayload, &pxJSON)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		randomSensor, err := db.GetRandomSensor(dbPostgres)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Generate PX Cookies
		cookies, err := px.GenCookie(pxJSON.TargetURL, pxJSON.URL, *randomSensor)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Encrypt the payload
		encryptedPayload, err := utils.Encrypt(cookies, utils.EncryptionKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Encode our encrypted payload in base64
		sEnc := base64.RawStdEncoding.EncodeToString([]byte(encryptedPayload))

		c.JSON(http.StatusOK, gin.H{"cookies": sEnc})
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
