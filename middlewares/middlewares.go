package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BasicAuthMiddleware(c *gin.Context) {
	errorENV := godotenv.Load()
	if errorENV != nil {
		log.Println("Error loading env file:", errorENV)
		panic("Failed to load env file")
	}
	// get the basic auth credentials
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized basic auth ERROR ~"})
		c.Abort()
		return
	}
	// check the credentials
	if username != os.Getenv("AUTH_USER") || password != os.Getenv("AUTH_PASS") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		c.Abort()
		return
	}
	c.Next()
}

func APIKeyAuthMiddleware(c *gin.Context) {
	errorENV := godotenv.Load()
	if errorENV != nil {
		log.Println("Error loading env file:", errorENV)
		panic("Failed to load env file")
	}
	log.Println("X-API-Key:", c.GetHeader("X-API-Key"))
	log.Println("AUTH_APIS:", os.Getenv("AUTH_API"))

	// get the api key
	apiKey := c.GetHeader("x-api-key")
	if apiKey == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized X-API-Key ERROR ~"})
		c.Abort()
		return
	}
	// check the api key

	if apiKey != os.Getenv("AUTH_APIS") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid X-API-Key"})
		c.Abort()
		return
	}
	c.Next()
}
