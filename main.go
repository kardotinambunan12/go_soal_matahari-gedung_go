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
	// "container/list"
	"fmt"
	// "os/user"
)

type Gedung struct {
	NamaGedung string
	Tinggi     int
}

type Gedungs []Gedung

func (u Gedungs) NameList() []string {
	var list []string
	for _, lbangunan := range u {
		if lbangunan.Tinggi > 3 {
			list = append(list, lbangunan.NamaGedung)
			

		}else{
			fmt.Println("Gedung yang tidak memenuhi syarat", lbangunan.NamaGedung)
		}
	}
	return list
}

func (i Gedungs)ListGedung()[]string{
	var lGedung []string
	for _, gedung := range i {
		if gedung.Tinggi >= 0{
			lGedung = append(lGedung, gedung.NamaGedung)
		}
	}
	return lGedung
}
func (i Gedungs)ListGedung2()[]string{
	var lGedung []string
	for _, gedung := range i {
		if gedung.Tinggi >= 9{
			lGedung = append(lGedung, gedung.NamaGedung)
		}else{
			fmt.Println("gedung yang tidak memnuhi syarat jika matahari 240 derajat : ", gedung.NamaGedung)
		}
	}
	return lGedung
}

func main() {
	gedungs := Gedungs{
		Gedung{NamaGedung: "GD1", Tinggi: 7},
		Gedung{NamaGedung: "GD2", Tinggi: 2},
		Gedung{NamaGedung: "GD3", Tinggi: 4},
		Gedung{NamaGedung: "GD4", Tinggi: 8},
		Gedung{NamaGedung: "GD5", Tinggi: 9},
	}

	UserList := gedungs.NameList()
	lgedung := gedungs.ListGedung()
	sgedung := gedungs.ListGedung2()

	fmt.Println("Gedung yang memenuhi syarat jika matahari tinggi 45 derajat:",UserList)
	fmt.Println("---------------------------------------------------")
	fmt.Println("Gedung yang memenuhi syarat jika sinar matahari tepat di atas :",lgedung)
	fmt.Println("---------------------------------------------------")
	fmt.Println("Gedung yang memenuhi syarat jika sinar matahari 240  derajat :",sgedung)
}
