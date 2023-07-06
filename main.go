//  soal :
//    clue nya matahari dan gedung 
//    dinama matahari menyinari sebuah gedung 
//    gedung paling depan 7 lantai, gedung belakangnya 2 lantai,
//    gedung belakangnya 4 lantai, gedung belakangnya 8 lantai,
//    gedung belakangnya 9 lantai,

//    pertanyaan ;;; gedung mana (berapa gedung)  yang terkena sinar matahari ?

// pertanyaan di bluebird  "btw ini belum tau benar atau enggak ya"
package main

import (
	"fmt"
)

type Gedung struct {
	NamaGedung string
	Tinggi     int
}

type Gedungs []Gedung

func (u Gedungs) FilterByTinggi(minTinggi int) []string {
	var list []string
	for _, gedung := range u {
		if gedung.Tinggi >= minTinggi {
			list = append(list, gedung.NamaGedung)
		}
	}
	return list
}

func main() {
	gedungs := Gedungs{
		Gedung{NamaGedung: "GD1", Tinggi: 7},
		Gedung{NamaGedung: "GD2", Tinggi: 2},
		Gedung{NamaGedung: "GD3", Tinggi: 4},
		Gedung{NamaGedung: "GD4", Tinggi: 8},
		Gedung{NamaGedung: "GD5", Tinggi: 9},
	}

	userList := gedungs.FilterByTinggi(3)
	lgedung := gedungs.FilterByTinggi(0)
	sgedung := gedungs.FilterByTinggi(9)

	fmt.Println("Gedung yang memenuhi syarat jika matahari tinggi 45 derajat:", userList)
	fmt.Println("---------------------------------------------------")
	fmt.Println("Gedung yang memenuhi syarat jika sinar matahari tepat di atas:", lgedung)
	fmt.Println("---------------------------------------------------")
	fmt.Println("Gedung yang memenuhi syarat jika sinar matahari 240 derajat:", sgedung)
}
