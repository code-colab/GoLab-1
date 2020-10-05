package cmd


import (
	"calculator/calc"
	"os"
	"fmt"
	"strconv"
)

func IsInt64(s string) bool {
	if _, err := strconv.ParseInt(s,10,64); err != nil {
		return false
	}
	return true
}

// Execute The function execute ...
func Execute() {
	// fmt.Println(os.Args[1:])
	// if !IsInt64(os.Args[2]) {
	// 	fmt.Println(calc.Operations[os.Args[2]])
	// }	

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