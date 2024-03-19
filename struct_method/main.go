package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Name  string
	Grade int
}

func (s Student) sayHello() {
	fmt.Println("Hello", s.Name)
}

func (s Student) getName(i int) string {
	return strings.Split(s.Name, " ")[i-1]
}
func main() {

	var s1 Student = Student{"Nichola Saputra", 100}

	var name = s1.getName(1)
	fmt.Println("Nama Panggilan : ", name)
}
