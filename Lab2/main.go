package main

import (
	"fmt"
)

var rec = Rectangle{
	a: 10,
	b: 20,
}

func main() {
	fmt.Println("Задание 2")
	Zadan2(5)
	fmt.Println("Задание 4")
	Zadan4("hello")
	fmt.Println("Задание 5")
	fmt.Println((CalcS(rec)))
	fmt.Println("Задание 6")
	fmt.Println(Zadan6(10, 5))
}

func Zadan2(num int) {

	if num < 0 {
		fmt.Println("Negative")
	} else if num > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Zero")
	}

}

func Zadan4(line string) {

	fmt.Println("длина строки ", line, " ", len(line))

}

type Rectangle struct {
	a int
	b int
}

func CalcS(rec Rectangle) int {

	return rec.a * rec.b

}
func Zadan6(num1 float32, num2 float32) float32 {

	return (num1 + num2) / 2

}
