//This example shows how to set and get cookies

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CookieTool() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Get cookie
		if cookie, err := c.Cookie("label"); err == nil {
			if cookie == "ok" {
				c.Next()
				return
			}
		}
		//cookie verification failed
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden with no cookie"})
		c.Abort()
	}
}

func main() {
	router := gin.Default()

	router.GET("/login", func(c *gin.Context) {
		//set cookie {"label": "ok"}, maxAge 30 seconds.
		c.SetCookie("label", "ok", 30, "/", "localhost", false, true)
		c.String(200, "login success")
	})

	router.GET("/home", CookieTool(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "Your home page"})
	})

	router.Run(":8080")
}
