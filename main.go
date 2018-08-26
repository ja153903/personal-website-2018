package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)


type BlogPost struct {
	BlogEntry string `json:"blog_post"`
	Date 	  string `json:"date"`
}


func handleTextPost(router *gin.Engine) {
	router.POST("/blog_entry", func(context *gin.Context) {
		var blogPost BlogPost

		context.BindJSON(&blogPost)
		
		context.JSON(200, gin.H{
			"status": "received",
			"blog_post": blogPost.BlogEntry,
			"date": blogPost.Date,
		})

		// at this point, we want to add the post to the table
	})
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	handleTextPost(router)

	router.Run(":8000")
}
