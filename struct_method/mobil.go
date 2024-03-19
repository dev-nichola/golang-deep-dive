package main

import "fmt"

type Mobil struct {
	Merek, Warna string
}

func (m Mobil) nyalakanMesin() {
	fmt.Println("Mesin Mobil", m.Merek, "manyala🔥🔥🔥")
}

func main() {
	var mobilToyota Mobil = Mobil{
		"Toyota", "Hitam",
	}

	mobilToyota.nyalakanMesin()
}
