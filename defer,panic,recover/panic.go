package main

import "fmt"

func division(num1, num2 int) int {
	if num2 == 0 {
		panic("Tidak Dapat Melakukan Pembagian")
	} else {
		return num1 / num2
	}
}

func Execute() string {
	return "Defer Akan Selalu Di Eksekusi"
}

func main() {

	defer fmt.Println(Execute())
	result := division

	fmt.Println(result(10, 0))
}
