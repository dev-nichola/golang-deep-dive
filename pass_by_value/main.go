package main

import "fmt"

func ubahNilai(x int) int {
	x = 100
	return x
}

func main() {
	x := 50
	fmt.Println(ubahNilai(x))
}
