package calc

import (
	"errors"
)

func Divide(a int64, b int64) (myReturn int64, err error) {
	err = nil
	if b == 0 {
		err = errors.New("Cannot divide by zero")
		return myReturn, err
	}

	myReturn = a / b
	return myReturn, err
}