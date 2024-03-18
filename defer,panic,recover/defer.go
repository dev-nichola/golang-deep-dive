package main

import "fmt"

func Execute() {
	fmt.Println("Program 4 Di Eksekusi")
}
func main() {

	// Pastikan Membuat Defer Itu Selalu Di Awal
	defer Execute() // Tetap Akan Di Eksekusi Walau Terjadi Error

	num := 0

	result := 10 / num
	fmt.Println(result) // Ini Bakalan Error

	fmt.Println("Program 1 Di Eksekusi")
	fmt.Println("Program 2 Di Eksekusi")
	fmt.Println("Program 3 Di Eksekusi")

}
