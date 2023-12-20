package main

import (
	"fmt"
	"math"
)

/*
Penerapan fungsi yang baik dapat menjadikan kode lebih modular dan dry(dont repeat yourself)
, tak perlu menuliskan banyak kode berulang kali, cukup sekali saja sesuai kebutuhan
*/

//func main() {
//	kataKataHarinIni("Tidak perlu kata-kata yang penting bukti nyata")
//}
//
//func kataKataHarinIni(message string) {
//	fmt.Println("Kata-kata hari ini :", message)
//}
//
//// deklarasi parameter bertipe sama : memberikan tipe data pada akhir param
//func funcA(param1, param2, param3 string) {
//
//}

//FUNGSI MULTIPLE RETURN
/*
Jika ada kebutuhan di
mana data yang dikembalikan harus banyak, biasanya digunakanlah tipe seperti
map , slice, atau struct sebagai nilai balik.
*/

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Println("area:", area)
	fmt.Println("circumference:", circumference)
}

func calculate(d float64) (float64, float64) {
	//hitung luas
	area := math.Pi * math.Pow(d/2, 2)

	//hitung keliling
	circumference := math.Pi * d

	return area, circumference
}
