package main 

import (
	"fmt"
	"strings"
)

type User struct {
	name string 
	age int64 
	password string
}
func change(age *int64){
    var a int64
    fmt.Scanf("%d",&a)
	*age = a
}

// func isOneIn(char string)bool{
// 	sumPlus := 0
// 	sumMinus := 0
//     for i := 0 ; i< len(char) ; i++ {
//         if char[i]=='+' {
// 			sumPlus++
// 		}else if char[i]=='-' {
// 			sumMinus++
// 		}
// 	}
// 	if (sumPlus<=1 && sumMinus==0) || (sumPlus==0 && sumMinus<=1) {
// 		return true
// 	}
// 	return false
// }


func goodFunc(char string){
	
	
}
func main(){
	a := strings.ReplaceAll("00009","","")
	fmt.Println(a)
}