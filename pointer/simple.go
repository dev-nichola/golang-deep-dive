package main

import "fmt"

func main() {
	name := "Nichola"
	name2 := &name

	*name2 = "Joko"

	fmt.Println(name)
}
