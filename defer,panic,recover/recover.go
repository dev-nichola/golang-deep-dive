package main

import "fmt"

func division(num1, num2 int) int {
	defer handlePanic()
	if num2 == 0 {
		panic("Tidak bisa membagi data")
	} else {
		return num1 / num2
	}
}

func handlePanic() {
	err := recover()

	if err != nil {
		fmt.Println("Recover", err)
	}
}

func main() {
	result := division

	fmt.Println(result(10, 0))
}
