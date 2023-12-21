package main

import "fmt"

/*
GoRoutine adalah sebuah cara untuk menjalankan fungsi secara konkuren
Eksekusi goroutine bersifat asynchronous, menjadikannya tidak saling tunggu dengan goroutine lain.
*/
//
//func print(till int, message string) {
//	for i := 0; i < till; i++ {
//		fmt.Println((i + 1), message)
//	}
//}
//
//func main() {
//	runtime.GOMAXPROCS(2)
//
//	go print(5, "Hallo") // "go" digunakan untuk membuat goroutine baru
//	print(5, "Apa kabar")
//
//	var input string
//	fmt.Scanln(&input)
//}

// PENGGUNAAN FUNGSI fmt.Scanln
// Fungsi ini akan meng-capture semua karakter sebelum user menekan tombol enter, lalu menyimpannya pada variabel.
func main() {
	var s1, s2, s3 string
	fmt.Scanln(&s1, &s2, &s3)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
