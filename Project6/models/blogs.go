package models



type Blogs struct {
	Id               int        `db:"id"`
	Title            string     `db:"title"`
	Body             string     `db:"body"`
	BlogerName       string     `db:"name"`
	BlogerLastname	 string     `db:"lastname"`
	Username         string     `db:"username"`
	Status           bool       `db:"status"`
	Created_at       string     `db:"created_at"`
	Updated_at       string     `db:"updated_at"`
	
}

type Blog struct {
	Title        string
	Body         string
	UserId       int 
}