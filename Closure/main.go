package main

import "fmt"

/*
Closure adalah fungsi yang bisa disimpan dalam sebuah variable
*/
//
//func main() {
//	var getMinMax = func(n []int) (int, int) {
//		var min, max int
//		for i, e := range n {
//			switch {
//			case i == 0:
//				max, min = e, e
//			case e > max:
//				max = e
//			case e < min:
//				min = e
//			}
//		}
//		return min, max
//	}
//	var numbers = []int{2, 3, 4, 5, 6, 7}
//	var min, max = getMinMax(numbers)
//	fmt.Printf("Data: %v\nmax : %v\nmin: %v\n", numbers, max, min)
//}

//Closure Immediately-Invoked Function Expression(IIFE)
//Close yang dieksekusi langsung pada saat deklarasinya. Digunakan untuk membungkus proses yang hanya sekali dilakukan

//func main() {
//	var numbers = []int{2, 3, 4, 5, 6, 1, 2, 5, 6, 7}
//
//	var filterNumber = func(min int) []int {
//		var r []int
//		for _, e := range numbers {
//			if e < min {
//				continue
//			}
//			r = append(r, e)
//		}
//		return r
//	}(4) //IIFE : adanya kurung paramter setelah fungsinya
//
//	fmt.Println("Original number:", numbers)
//	fmt.Println("Filtered number:", filterNumber)
//}

//Closure sebagai nilai return

func main() {
	numbers := []int{2, 3, 4, 5, 6, 2, 1, 2, 4, 1, 2, 2, 8, 4}
	max := 2

	var howMany, numbersList = findMax(numbers, max)
	var theNumbers = numbersList()

	fmt.Println("Found : ", howMany)
	fmt.Println("Numbers : ", theNumbers)
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}
