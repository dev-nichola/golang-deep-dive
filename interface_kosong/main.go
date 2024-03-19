package main

import "fmt"

func Cetak(nilai any) {
	fmt.Println("Nilai : ", nilai)
}

func main() {
	Cetak("Hello Golang!")
	Cetak(100)
	Cetak(true)
}
