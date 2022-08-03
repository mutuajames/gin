package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//Load the templates
	router.LoadHTMLGlob("templates/*")
	//define the route handler
	//router.GET("/", func(context *gin.Context) {
	//	//call the html method of the context to render a template
	//	context.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page"})
	//})
	router.GET("/", showIndexPage)
	router.GET("article/view/:article_id", getArticle)
	router.Run()
}
