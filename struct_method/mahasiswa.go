package main

import "fmt"

type Karyawan struct {
	Nama, Alamat, Jabatan string
	Gaji                  int
}

func (k Karyawan) Construct() {
	fmt.Println("Saya", k.Nama, "Seorang", k.Jabatan, "Rumah Saya Ada Di", k.Alamat, "Gaji Saya Sekitar", k.Gaji, "Juta Rupiah")
}

func main() {
	var karyawan Karyawan = Karyawan{
		"Nichola Saputra",
		"Blitar",
		"Programmer",
		10,
	}

	karyawan.Construct()
}
