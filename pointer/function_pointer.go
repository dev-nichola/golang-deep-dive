package main

import "fmt"

func main() {
	name := "Nichola"
	changeName(&name)
	fmt.Println(name)
}

func changeName(name *string) {
	*name = "Joko Moro"
	fmt.Println(*name)
}
