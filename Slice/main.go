package main

import "fmt"

/*
slice bisa dibilang array tapi tidak memiliki panjang yang spesifik.
*/

func main() {
	//var animalsA = []string{"lion,cat,zebra"}     //slice
	//var animalsB = [4]string{"bull", "whale", "shark", "cat"} //array
	//var animalsC = [...]string{"bull,whale,shark"} //array

	//membuat slice dari sebuah array
	//var newAnimals = animalsB[0:3] //mengakses array animalsB mulai dari indeks 0 sampai indeks sebelum 3
	//fmt.Println(newAnimals)

	//slice merupakan tipe data reference
	//var animals = []string{"dog", "hyena", "bull", "shark"}
	//
	//Aanimals := animals[0:3]
	//Banimals := animals[2:4]
	//Canimals := animals[0:2]
	//
	//fmt.Println(animals)
	//fmt.Println(Aanimals)
	//fmt.Println(Banimals)
	//fmt.Println(Canimals)
	//
	//Canimals[1] = "fish"
	//
	//fmt.Println(animals)
	//fmt.Println(Aanimals)
	//fmt.Println(Banimals)
	//fmt.Println(Canimals)
	//semua berubah karena slice merupakan tipe data reference

	//fungsi len : menghitung jumlah elemen slice yang ada.
	//animals := []string{"cat", "dog", "hyena"}
	//fmt.Println(len(animals)) //3

	//fungsi cap : untuk menghitung lebar atau kapasitas maksimum
	//animals := []string{"cat", "dog", "hyena"}
	//fmt.Println(len(animals)) //3
	//fmt.Println(cap(animals)) //3
	//
	//animalsA := animals[0:2]
	//fmt.Println(len(animalsA)) //2
	//fmt.Println(cap(animalsA)) //3
	//
	//animalsB := animals[1:2]
	//fmt.Println(len(animalsB)) //1
	//fmt.Println(cap(animalsB)) //2

	//fungsi append : digunakan untuk menambahkan elemen pada elemen terakhir
	//animals := []string{"cat", "dog", "hyena"}
	//animalsA := append(animals, "shark")
	//
	//fmt.Println(animals)
	//fmt.Println(animalsA)

	//fungsi copy : digunakan untuk men-copy elements slice pada src (parameter ke-2), ke dst (parameter pertama)
	//dst := make([]string, 3)
	//src := []string{"cat", "dog", "shark", "hyena"}
	//n := copy(dst, src)
	//
	//fmt.Println(n)

	//Pengaksesan elemen slice dengan 3 indeks, indeks ke-3 berarti kapasitas slice
	animals := []string{"cat", "dog", "lion", "hyena"}
	animalsA := animals[0:2:2]

	fmt.Println(animals)
	fmt.Println(len(animals))
	fmt.Println(animalsA)
	fmt.Println(len(animalsA))
}
