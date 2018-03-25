// Package main is the CLI.
// You can use the CLI via Terminal.
package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/madhums/go-gin-mgo-demo/db"
	"github.com/madhums/go-gin-mgo-demo/gin_html_render"
	"github.com/madhums/go-gin-mgo-demo/handlers/articles"
	"github.com/madhums/go-gin-mgo-demo/middlewares"
)

const (
	// Port at which the server starts listening
	Port = "7000"
)

func init() {
	db.Connect()
}

func main() {

	// Configure
	router := gin.Default()

	// Set html render options
	htmlRender := GinHTMLRender.New()
	htmlRender.Debug = gin.IsDebugging()
	htmlRender.Layout = "basic/main"
	// htmlRender.TemplatesDir = "templates/" // default
	// htmlRender.Ext = ".html"               // default

	// Tell gin to use our html render
	router.HTMLRender = htmlRender.Create()

	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = false

	// Middlewares
	router.Use(middlewares.Connect)
	router.Use(middlewares.ErrorHandler)

	// Statics
	router.Static("/public", "./public")

	// Routes

	router.GET("/", func(c *gin.Context) {

<<<<<<< HEAD
		c.Redirect(http.StatusMovedPermanently, "/main")
	})

	router.GET("/doctors", func(c *gin.Context) { // 의사소개
		c.HTML(http.StatusOK, "basic/doctors", gin.H{})
=======
	router.GET("/survey", func(c *gin.Context) {
		c.HTML(http.StatusOK, "basic/survey", gin.H{})
>>>>>>> 131ced1d3bffd50363b92fcd99c5ad2d153bf7ed
	})

	router.GET("/main", func(c *gin.Context) { //메인페이지
		c.HTML(http.StatusOK, "basic/content", gin.H{})
	})

	router.GET("/contact", func(c *gin.Context) { //메인페이지
		c.HTML(http.StatusOK, "basic/contact", gin.H{})
	})

	// Articles
	router.GET("/new", articles.New)                     // 새로운 메일 보내기
	router.GET("/consult/:_id", articles.Edit)           // 메일 읽기
	router.GET("/consults", articles.List)               // 내가 쓴 상담 리스트
	router.POST("/consult", articles.Create)             // 상담 쓰기
	router.POST("/consult/:_id", articles.Update)        // 상담 업데이트
	router.POST("/delete/consult/:_id", articles.Delete) // 상담 지우기

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
}
