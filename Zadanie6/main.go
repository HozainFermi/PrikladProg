package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	num1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	num2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	num3, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d", (num1+num2+num3)/3)

}
