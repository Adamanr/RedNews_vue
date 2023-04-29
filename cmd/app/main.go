package main

import (
	"RedNews_Server/internal/config"
	https "RedNews_Server/internal/controller/http"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}

		c.Next()
	}
}
func main() {
	config.ReadConfig()

	router := gin.New()
	router.Use(CorsMiddleware())
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})
	https.Router(router)
	if err := router.Run("localhost:8080"); err != nil {
		log.Panicln(err)
		return
	}
}
