package main

import "fmt"

type Mahasiswa struct {
	Name, Address string
	Age           int
}

func main() {
	var mahasiswa Mahasiswa = Mahasiswa{
		Name:    "Nichola",
		Address: "Blitar",
		Age:     20,
	}

	name := mahasiswa.Name
	address := mahasiswa.Address
	age := mahasiswa.Age

	fmt.Println(name)
	fmt.Println(address)
	fmt.Println(age)

}
