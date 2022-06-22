package models

type Person struct {
	Name string
	Family string 
}
type Book struct {
	Id int
	Name string
	Year uint32
	Author Person 
}

var Books map[int]Book 

