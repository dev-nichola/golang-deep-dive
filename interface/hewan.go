package main

import "fmt"

type Hewan interface {
	Bersuara() string
}

type Anjing struct{}

func (a Anjing) Bersuara() string {
	return "Guk Guk"
}

type Kucing struct{}

func (k Kucing) Bersuara() string {
	return "Meow Meow"
}

func main() {
	anjing := Anjing{}
	kucing := Kucing{}

	var hewan Hewan = anjing
	fmt.Println(hewan.Bersuara())

	hewan = kucing
	fmt.Println(hewan.Bersuara())
}
