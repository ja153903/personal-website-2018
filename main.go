package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"database/sql"
	"context"
	_ "github.com/go-sql-driver/mysql"
)

var ctx = context.Background()

type BlogPost struct {
	BlogEntry string `json:"blog_post"`
	Date 	  string `json:"date"`
}

func getTextPost(router *gin.Engine) {
	router.GET("/api/blog_entry/get", func(context *gin.Context) {
		blogPosts := map[string]string{}

		db, _ := sql.Open("mysql", "root:password@/blog")
		defer db.Close()

		sqlQuery := `
		SELECT * FROM POSTS;
		`
		rows, err:= db.QueryContext(ctx, sqlQuery)
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			var blogEntry string
			var datePosted string 

			if err := rows.Scan(&blogEntry, &datePosted); err != nil {
				panic(err)
			}

			blogPosts[datePosted] = blogEntry 
		}
		
		context.JSON(200, gin.H{
			"status": "completed",
			"data": blogPosts,
		})
	})
}

func handleTextPost(router *gin.Engine) {
	router.POST("/api/blog_entry/post", func(context *gin.Context) {
		var blogPost BlogPost

		context.BindJSON(&blogPost)
		
		context.JSON(200, gin.H{
			"status": "received",
			"blog_post": blogPost.BlogEntry,
			"date": blogPost.Date,
		})

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
	getTextPost(router)

	router.Run(":8000")
}
