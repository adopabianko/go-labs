package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Example to test
//  curl --location 'http://localhost:8080/test' \
// --header 'X-App-Token: 123456' \
// --header 'X-Client-Id: service-name'

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		xAppToken := len(c.GetHeader("X-App-Token"))
		xClientID := len(c.GetHeader("X-Client-Id"))

		if xAppToken == 0 || xClientID == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "header X-App-Token or X-Client-Id is mandatory",
				"data":    nil,
			})
			return
		}

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println("----", status)
		fmt.Println("") // new line
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
