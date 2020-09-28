package cmd


import (
	"calculator/calc"
	"os"
	"fmt"
)

// Execute The function execute ...
func Execute() {

	// primeironumero, _ := strconv.ParseInt(os.Args[1], 10, 64)
	// segundonumero, _ := strconv.ParseInt(os.Args[2], 10, 64)
	// var myVariable int64 = 6
	// myReturn, _ := calc.Sum(primeironumero,segundonumero)
	//  fmt.Println(myReturn)
	// myReturn, _ = calc.Divide(myVariable, 0)
	// fmt.Println(myReturn)

	o, err := calc.CheckOperation("+")
	if err != nil {
		os.Exit(127)
	}

	r, _ := o.Calculate()
	fmt.Println(r)


	o, err = calc.CheckOperation("-")
	if err != nil {
		os.Exit(127)
	}

	r, _ = o.Calculate()
	fmt.Println(r)

	o, err = calc.CheckOperation("*")
	if err != nil {
		os.Exit(127)
	}

	r, _ = o.Calculate()
	fmt.Println(r)

	o, err = calc.CheckOperation("/")
	if err != nil {
		os.Exit(127)
	}

	r, _ = o.Calculate()
	fmt.Println(r)
}