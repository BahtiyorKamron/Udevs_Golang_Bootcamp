package main

import (
	// "net/http"
    "Project6/controllers"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
    
	router.POST("/signup",controllers.SignUp)
	router.POST("/signin",controllers.SignIn)
	router.POST("/post-blog",controllers.PostBlog)
	router.GET("/blogs/:id",controllers.GetBlogs)
	router.Run()
}