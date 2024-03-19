package main

import "fmt"

type MataKuliah struct {
	Nama string
}
type Mahasiswa interface {
	Nama() string
	NIM() string
	AmbilKuliah(mk MataKuliah) string
}

type S1 struct {
	nama, nim string
}

func (s1 S1) Nama() string {
	return s1.nama
}

func (s1 S1) NIM() string {
	return s1.nim
}

func (s1 S1) AmbilKuliah(mk MataKuliah) string {
	return fmt.Sprintf("%s - %s - %s", s1.NIM(), s1.nama, mk.Nama)
}

type S2 struct {
	nama, nim string
}

func (s2 S2) Nama() string {
	return s2.nama
}

func (s2 S2) NIM() string {
	return s2.nim
}

func (s2 S2) AmbilKuliah(mk MataKuliah) string {
	return fmt.Sprintf("%s - %s - %s", s2.NIM(), s2.nama, mk.Nama)
}

func main() {
	mk := MataKuliah{Nama: "Pemrogramana Golang"}
	mahasiswaS1 := S1{"Nichola", "11223344"}

	var mhs Mahasiswa = mahasiswaS1
	fmt.Println(mhs.AmbilKuliah(mk))

	mk = MataKuliah{"Pemograman TypeScript"}
	mahasiswaS2 := S2{"Joko", "11223344"}
	mhs = mahasiswaS2
	fmt.Println(mhs.AmbilKuliah(mk))
}
