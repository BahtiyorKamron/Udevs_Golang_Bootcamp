package bigint

import (
	"errors"
    "strconv"
	"strings"
	
)

type Bigint struct {
	value string
}

func isOneIn(char string)bool{
	sumPlus := 0
	sumMinus := 0
    for i := 0 ; i< len(char) ; i++ {
        if char[i]=='+' {
			sumPlus++
		}else if char[i]=='-' {
			sumMinus++
		}
	}
	
	if (sumPlus<=1 && sumMinus==0) || (sumPlus==0 && sumMinus<=1) {
		return true
	}
	return false
}

func validate(num *string)error{

	str := *num

    

	if (strings.Contains(str,"-") || strings.Contains(str,"+") ) &&  string(str[0])!="-" && string(str[0])!="+"  {
		return errors.New("Bad Input")
	}
	
	
	allowed := "+-0123456789"
	for i := 0 ; i<len(str) ; i++ {
		if !strings.Contains(allowed,string(str[i])) {
			return errors.New("Bad Input")
		}
		
	}


	if isOneIn(str)!=true{
		return errors.New("Bad Input")
	}
	
	


	removingZero := 0
	var newNum string
	tire := ""
	if string(str[0]) == "-" {
        tire = "-"
	}
	str = strings.ReplaceAll(str,"-","")
	str = strings.ReplaceAll(str,"+","")
	for j := 0 ; j < len(str) ; j++ {
        if(string(str[j]) != "0"){
			removingZero = 1
			
		}
		if(removingZero==1){
			newNum += string(str[j])
		}

	}
	
	
	*num = tire + newNum
	
	return nil
}

func NewInt(num string)(Bigint,error){
	err :=  validate(&num)
	if err!=nil {
		return Bigint{},err
	}
	return Bigint{ value:num } , nil
}

func (z *Bigint) Set(num string)error{
	err :=  validate(&num)
	if err!=nil {
		return err
	}
	z.value = num 
	return nil
}

func Add(a,b Bigint) Bigint {
	d,_ := strconv.Atoi(a.value)
	c,_ := strconv.Atoi(b.value)
	
	return Bigint{
		value : strconv.Itoa(d+c) ,
	}
}
func Multiply(a,b Bigint) Bigint {
	d,_ := strconv.Atoi(a.value)
	c,_ := strconv.Atoi(b.value)
	
	return Bigint{
		value : strconv.Itoa(d*c) ,
	}
}
func Mod(a,b Bigint) Bigint {
	d,_ := strconv.Atoi(a.value)
	c,_ := strconv.Atoi(b.value)
	
	return Bigint{
		value : strconv.Itoa(d%c) ,
	}
}