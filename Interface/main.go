package main

import (
	"fmt"
	"math"
)

/*
Interface adalah kumpulan definisi method yang tidak memiliki isi
*/

//type hitung interface {
//	luas() float64
//	keliling() float64
//}
//
//type lingkaran struct {
//	diameter float64
//}
//
//func (l lingkaran) jariJari() float64 {
//	return l.diameter / 2
//}
//
//func (l lingkaran) luas() float64 {
//	return math.Pi * math.Pow(l.jariJari(), 2)
//}
//
//func (l lingkaran) keliling() float64 {
//	return math.Pi * l.diameter
//}
//
//func main() {
//	var l1 = lingkaran{diameter: 20}
//	fmt.Println("Jari-jari :", l1.jariJari())
//	fmt.Println("Luas :", l1.luas())
//	fmt.Println("Keliling :", l1.keliling())
//}

// EMBEDED INTERFACE : INTERFACE BISA DIEMBEDED KE INTERFACE LAIN
type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume(float642 float64)
}

type hitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k kubus) keliling() float64 {
	return k.sisi * 12
}

func (k kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func main() {
	var k1 = kubus{sisi: 6}
	fmt.Println("Sisi : ", k1.sisi)
	fmt.Println("Volume : ", k1.volume())
	fmt.Println("Luas : ", k1.luas())
	fmt.Println("Keliling : ", k1.keliling())
}
