package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()

	// CORS for all origins
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	xml, err := os.ReadFile("sample.xml")
	if err != nil {
		log.Fatal(err)
	}

	// VAST XMLレスポンスを返却するエンドポイント
	r.GET("/vast", func(c *gin.Context) {
		vastXML := xml
		c.Data(http.StatusOK, "application/xml; charset=utf-8", []byte(vastXML))
	})

	r.Run(":8080")
}
