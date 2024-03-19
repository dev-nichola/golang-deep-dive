package main

import "fmt"

type Cat interface {
	setName(name string) string
	setMeow() string
}

type MyCat struct{}

func (c MyCat) setName(name string) string {
	return "Hi, " + name
}

func (c MyCat) setMeow() string {
	return c.setName("Nichola") + "! Meoww..."
}

func callName(cat Cat) string {
	return cat.setMeow()
}

func main() {
	var myCat MyCat
	fmt.Println(callName(myCat))
}
