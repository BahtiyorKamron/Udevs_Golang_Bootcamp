package database

import (
    // "database/sql"
    "fmt"
    "log"
    "Project6/models"
    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)



func GetUsers(query string) ([]models.User,  error ){
	db ,err := sqlx.Connect("postgres","host=localhost dbname=newdb user=kamron password=kamron sslmode=disable")

	if err != nil {
		log.Fatalln(err)
	}

	result := []models.User{}
	err = db.Select(&result,query)
    
	return result,err

}

func UserByUsername(query string,username string)(bool,error){
	db ,err := sqlx.Connect("postgres","host=localhost dbname=newdb user=kamron password=kamron sslmode=disable")
    
	if err != nil {
		log.Fatalln(err)
	}
	result := []models.User{}
	err = db.Select(&result,query,username)
	if len(result) > 0 {
		return true , err 
	}else {
		return false , err 
	}
}

func GetUserByUsername(query string,username,password string)(bool , error){
	db ,err := sqlx.Connect("postgres","host=localhost dbname=newdb user=kamron password=kamron sslmode=disable")
    
	if err != nil {
		log.Fatalln(err)
	}
    fmt.Println(username,password)
	result := []models.User{}
	err = db.Select(&result,query,username,password)
	fmt.Println(result)

	if len(result) > 0 {
		return true , err 
	}else {
		return false , err 
	}
	

}
func InsertUser(query string,user models.User){
	db ,err := sqlx.Connect("postgres","host=localhost dbname=newdb user=kamron password=kamron sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
    
	db.MustExec(query,user.Username,user.Name,user.Lastname,user.Password)

	

}
func InsertBlog(query string,blog models.Blog){
	db ,err := sqlx.Connect("postgres","host=localhost dbname=newdb user=kamron password=kamron sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(query,blog.Title,blog.Body,blog.UserId)
}
func GetBlogs(query string,id string)([]models.Blogs,error){
	db ,err := sqlx.Connect("postgres","host=localhost dbname=newdb user=kamron password=kamron sslmode=disable")

	if err != nil {
		log.Fatalln(err)
	}

	blogs := []models.Blogs{}
	err = db.Select(&blogs,query,id)

	return blogs , err 
}

// func UpdateUser(query string ) ([]models.User , error){

    //  return nil
	// db.MustExec(`with old_data as (
	// 	select
	// 		id,
	// 		username,
	// 		password,
	// 		name,
	// 		lastname
	// 	from 
	// 		users
	// 	where id = $1 
	// )update users as u set
	// 	username = (
	// 		case
	// 			when length($2) > 1 then $2
	// 			else o.username
	// 		end
	// 	),
	// 	password = (
	// 		case
	// 			when length($3) > 1 then crypt($3, gen_salt('bf', 8))
	// 			else o.password
	// 		end
	// 	),
	// 	name = (
	// 	  case
	// 		   when length($4) > 1 then $4
	// 		   else o.name
	// 	  end
	// 	),
	// 	lastname = (
	// 	    case 
	// 		   when length($5) > 1 then $5
	// 		   else o.lastname 
	// 		end
	// 	)
	// from old_data as o
	// where u.id = $1
	// returning u.*   `,"1","Dostonbek","asdfsadf","doston11","Bozorboyev")
// }