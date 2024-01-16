package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define a route that matches any HTTP method for any path
	r.Any("/*path", func(c *gin.Context) {
		method := c.Request.Method
		path := c.Param("path")
		url := fmt.Sprintf("https://nasainsbury.com%s", path)

		req, err := http.NewRequest(method, url, c.Request.Body)

		if err != nil {
			message := fmt.Sprintf("Unable to make request.\nMethod: %s.\nURL: %s", method, url)
			c.JSON(200, gin.H{
				"url":     url,
				"method":  method,
				"message": message,
			})
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			message := fmt.Sprintf("Unable to make request.\nMethod: %s.\nURL: %s", method, url)
			c.JSON(200, gin.H{
				"url":     url,
				"method":  method,
				"message": message,
			})
			return
		}

		defer resp.Body.Close()

		c.JSON(200, gin.H{
			"url":    url,
			"method": method,
			"response": gin.H{
				"status": resp.Status,
			},
		})
	})

	r.Run(":8080")
}
