package controllers

import (
	"net/http"
    "Project6/models"
	"Project6/database"
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
)

func PostBlog(c *gin.Context){
    var blog models.Blog 

	err := c.ShouldBindJSON(&blog)

	if err != nil {
		log.Fatalln(err)
	}

	database.InsertBlog("insert into blogs(title,body,user_id,created_at,status) values($1,$2,$3,Now(),true)",blog)
	c.JSON(201,gin.H{"error":"Blog has posted"})
	// (query,blog.Title,blog.Body,blog.BlogerName,blog.BlogerLastname,blog.UserByUsername)
}
func GetBlogs(c *gin.Context){
	// var blogs models.Blogs
	fmt.Println(c.Param("id"))
	datas , err := database.GetBlogs(`SELECT 
	b.id,
	b.title,
	b.created_at,
	b.body,
	u.name,
	u.lastname
 FROM blogs b 
 LEFT JOIN users u ON b.user_id=u.id WHERE u.id=$1`,c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK,datas)
}