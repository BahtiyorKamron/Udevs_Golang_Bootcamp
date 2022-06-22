package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"BookShelf/models"
	"fmt"
)




var Books map[int]models.Book


func init(){
	
	Books = make(map[int]models.Book)
		Books[1] = models.Book{
			Id: 1,
			Name: "Animal Farm",
			Year: 1945,
			Author: models.Person{
				Name: "Jorj",
				Family: "Oruel",
			},
		}
		Books[2] = models.Book{
			Id: 2,
			Name: "Raqamli qal'a",
			Year: 1997,
			Author: models.Person{
				Name: "Dan",
				Family: "Braun",
			},
		}
}

func GetBooks(c *gin.Context) {


	c.IndentedJSON(http.StatusOK, Books)
	
}

func PostBook(c *gin.Context){
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors":err.Error()})
		return
	}
    fmt.Println(Books)
	if _,ok := Books[book.Id]; ok {
		c.JSON(http.StatusBadRequest, gin.H{"errors":"Bu ID dagi kitob mavjud"})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"message":"Succesfully added new book",
		"data":book,
	})
}

func UpdateBook(c *gin.Context){
	var book models.Book
	if err := c.ShouldBindJSON(&book) ; err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	if _,ok := Books[book.Id]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"errors":"Bunday kitob mavjud emas"})
		return
	}
	Books[book.Id]=book
	c.JSON(http.StatusOK,gin.H{"succes changed the book":book})

}

func DeleteBook(c *gin.Context){
	var book models.Book 
	if err := c.ShouldBindJSON(&book) ; err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	if _,ok := Books[book.Id]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"errors":"Bunday kitob mavjud emas"})
		return
	}
	delete(Books,book.Id)
	c.JSON(202,gin.H{"Message":"Succesfully deleted"})
}