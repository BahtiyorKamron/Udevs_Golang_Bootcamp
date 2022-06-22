package controllers

import (
	"net/http"
    "Project6/models"
	"Project6/database"
	"github.com/gin-gonic/gin"

	"log"
)



func SignUp(c *gin.Context){
    var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	    return
	}


	userExists , err := database.UserByUsername("select * from users where username=$1",user.Username)
	if err != nil {
		log.Fatalln(err)
		return
	}

	if userExists {
		c.JSON(http.StatusBadRequest,gin.H{"error":"This username is already in use"})
		return
	}
    
	database.InsertUser("insert into users(username,name,lastname,password) values($1,$2,$3,crypt($4, gen_salt('bf', 8))) returning *",user)
    

	c.JSON(201,gin.H{"status":"You have succesfully registired"})

}
func SignIn(c *gin.Context){
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		log.Fatalln(err)
		return
	}
	if (user.Username=="" || user.Password=="")  {
		c.JSON(http.StatusBadRequest,gin.H{"error":"You need enter login and password"})
		return
	}
	currentUser , err := database.GetUserByUsername("select * from users where username=$1 and password=crypt($2, password)",user.Username,user.Password)
	if !currentUser {
		c.JSON(404,gin.H{"error":"That user does not exist"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"You have succcesfuly signed in"})
}

