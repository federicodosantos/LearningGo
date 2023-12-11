package main

import "fmt"

/*
Array adalah kumpulan data bertipe sama, yang disimpan dalam sebuah variable
*/

func main() {
	var hewan = [4]string{"anjing", "badak", "iguana", "hyena"}
	fmt.Println(hewan[0])
	fmt.Println(hewan[1])
	fmt.Println(hewan[2])
	fmt.Println(hewan[3])

	//insialisasi array tanpa jumlah elemen
	var buah = [...]string{"jeruk", "mangga", "apel"}
	fmt.Println("Isi array: ", buah)
	fmt.Println("Panjang array: ", len(buah))

	//array multidimensi
	var numbers = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Isi array: ", numbers)
	fmt.Println("Panjang array: ", len(numbers))

}
