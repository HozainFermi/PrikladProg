package Lab1

import "fmt"

func Calc(operation string, firstnumber float64, secondnumber float64) {
	if operation == "+" {
		fmt.Printf("%f", firstnumber+secondnumber)
	} else if operation == "-" {
		fmt.Printf("%f", firstnumber-secondnumber)
	} else {
		fmt.Printf("Только - или +")
	}

}
