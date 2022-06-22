package main 

import (
	"github.com/gin-gonic/gin"
	// "net/http"
	"BookShelf/controllers"
)

func main(){
	router  := gin.Default()
	v1 := router.Group("v1")
	v1.GET("/books",controllers.GetBooks)
	v1.POST("/books",controllers.PostBook)
	v1.PUT("/books",controllers.UpdateBook)
	v1.DELETE("/books",controllers.DeleteBook)
	router.Run()
}