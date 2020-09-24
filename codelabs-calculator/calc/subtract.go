package calc

func Subtract(a int64, b int64) (myReturn int64, err error) {
	myReturn = (a - b)
	return myReturn, nil
}
