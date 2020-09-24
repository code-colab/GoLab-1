package calc

func Multiply(a int64, b int64) (myReturn int64, err error) {
	myReturn = (a * b)
	return myReturn, nil
}
