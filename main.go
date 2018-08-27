package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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
		db, err := sql.Open("mysql", "root:password@/blog")
		if err != nil {
			panic(err)
		}

		defer db.Close()
		
		sqlQuery := `
		INSERT INTO POSTS (blog_post, date_posted)
		VALUES (?, ?)
		`
		_, err = db.Exec(sqlQuery, blogPost.BlogEntry, blogPost.Date)
		if err != nil {
			panic(err)
		}

	})
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	handleTextPost(router)

	router.Run(":8000")
}
