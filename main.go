package main

import (
	"Pengemudi/models"
	//"day1/models"
	"fmt"
)

func main() {

	var pgm models.Pengemudi
	pgm.NIP = "3884"
	pgm.Nama = "FARHAN"
	pgm.Alamat = "MAMPANG"

	type (
		Student struct {
			NIM string `json:"NIM"`
		}
	)

	var std Student
	std.NIM = "FARHAN"

	fmt.Println(pgm.NIP + " " + pgm.Nama + " " + pgm.Alamat)

	var armada models.Kendaraan
	armada.NoMobil = "B1PR"
	armada.Pool = "KJ"

	fmt.Print(" ", armada.NoMobil, " ", armada.Pool)

}
