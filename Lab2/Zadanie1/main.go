package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	if num%2 == 0 {
		fmt.Println("Введеное число - четное")
	} else {
		fmt.Println("Введеное число - нечетное")
	}

}
