package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)


func handleTextPost(router *gin.Engine) {
	router.POST("/blog_post", func(context *gin.Context) {
		blog_post := context.PostForm("blog_post")
		context.JSON(200, gin.H{
			"status": "received",
			"blog_post": blog_post,
		})
	})
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	handleTextPost(router)

	router.Run(":8000")
}
