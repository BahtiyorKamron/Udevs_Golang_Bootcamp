package models

type User struct {
	Id        int      `db:"id"`
	Name      string   `db:"name"`
	Lastname  string   `db:"lastname"`
	Username  string   `db:"username"`
	Password  string   `db:"password"`
}