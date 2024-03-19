package main

import "fmt"

type Items struct {
	NamaBarang string
	Kuantitas  int64
}
type Keranjang struct {
	NamaPemilik string
	Items       []Items
	Total       float64
}

func newKeranjang(nama string) Keranjang {
	return Keranjang{
		NamaPemilik: nama,
		Items:       []Items{},
		Total:       0,
	}
}

func main() {
	keranjang1 := newKeranjang("Keranjang Belanja CPM")

	fmt.Println(keranjang1)

	keranjang2 := newKeranjang("Keranjang Belanja 2")

	fmt.Println(keranjang2)
}
