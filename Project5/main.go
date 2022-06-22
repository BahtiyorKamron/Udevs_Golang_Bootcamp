package main 

import (
	"fmt"
	"bigint/bigint/bigint"
)

func main(){
    a , err := bigint.NewInt("+432")
	if err != nil {
		panic(err)
	}

	err = a.Set("2")
	if err != nil {
		panic(err)
	}

	b , err := bigint.NewInt("1")
	if err != nil {
		panic(err)
	}
	// c:= bigint.Add(a,b)
	// c := bigint.Multiply(a,b)
	c := bigint.Mod(a,b)

	fmt.Println(c)
	
	
}