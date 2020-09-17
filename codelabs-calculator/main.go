package main

import (
	"fmt"
	_ "reflect"
	"errors"
	_ "log"
)

func sum(a int64, b int64) (myReturn int64, err error) {
	myReturn = (a + b)
	return myReturn, nil
}

func divide(a int64, b int64) (myReturn int64, err error) {
	err = nil
	if b == 0 {
		err = errors.New("Cannot divide by zero")
		return myReturn, err
	}

	myReturn = a / b
	return myReturn, err
}

func subtract(a int64, b int64) (myReturn int64, err error) {
	myReturn = (a - b)
	return myReturn, nil
}

func multiply(a int64, b int64) (myReturn int64, err error) {
	myReturn = (a * b)
	return myReturn, nil
}

func main() {
	var myVariable int64 = 6
	myReturn, _ := sum(myVariable, myVariable)
	fmt.Println(myReturn)
	myReturn, _ = divide(myVariable, 0)
	fmt.Println(myReturn)
}


