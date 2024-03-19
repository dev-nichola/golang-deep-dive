package main

import "fmt"

type Mahasiswa struct {
	Name, NIM string
}

func main() {
	var mhs Mahasiswa = Mahasiswa{"Nichola", "1223344"}

	// memberitahu bahwa mhs1 refference
	mhs1 := &mhs

	fmt.Println(mhs1)

	// refference
	*mhs1 = Mahasiswa{"Joko", "Budi"}

	fmt.Println(mhs)
}
