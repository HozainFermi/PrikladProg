package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var sign string

	flag.StringVar(&sign, "o", "operation", "Specify operation")

	flag.Parse()
	num1, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	num2, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		panic(err)
	}

	switch sign {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		minus := num1 - num2
		fmt.Println(minus)
	case "/":
		fmt.Println(num1 / num2)
	case "%":
		fmt.Println(num1 % num2)

	}
	fmt.Println("кол-во арг - ", len(os.Args))

}
