package main

import (
	"fmt"
	"strings"
)

/*
Variadic function atau pembuatan fungsi dengan parameter sejenis yang tak terbatas.
Jadi parameter yang disisipkan bisa berapa saja.
*/

//func main() {
//	//avg := calculate(2, 3, 1, 2, 3, 4, 5, 6, 7, 8)
//	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //menggunakan slice
//	avg := calculate(numbers...)
//	msg := fmt.Sprintf("Rata-rata:%.2f", avg)
//	fmt.Println(msg)
//}
//
//func calculate(numbers ...int) float64 {
//	total := 0
//
//	for _, number := range numbers {
//		total += number
//	}
//
//	avg := float64(total) / float64(len(numbers))
//
//	return avg
//}

// Penggunaan parameter biasa dan variadic di fungsi
func main() {
	yourHobbies("Riko", "Swimming,Running,Gamming")
}

func yourHobbies(name string, hobbies ...string) {
	hobbiesAsstring := strings.Join(hobbies, ",")

	fmt.Printf("Hello my name is:%s\n", name)
	fmt.Printf("My hobbies are:%s\n", hobbiesAsstring)
}
