package main

import (
	"fmt"
	_ "reflect"
	"errors"
	_ "log"
)

func Sum(a int64, b int64) (int64, error) {
	myReturn := (a + b)
	return myReturn, nil
}

func Divide(a int64, b int64) (myReturn int64, err error) {
	err = nil
	if b == 0 {
		err = errors.New("Cannot divide by zero")
		return myReturn,err
	}

	myReturn = a / b
	return myReturn, err
}

func main() {

	
	var myVariable int64 = 6
	myReturn,_ := Sum(myVariable, myVariable)
	fmt.Println(myReturn)
	myReturn, _ = Divide(myVariable, 0)
	fmt.Println(myReturn)
}


