package main

import "fmt"

/*
Map adalah tipe data asosiatif, berbentuk key-value pair
*/

func main() {
	//chicken := map[string]int{} // tipe data key = string, tipe data value = integer
	//
	//chicken["january"] = 20
	//chicken["february"] = 10
	//
	//fmt.Println("january", chicken["january"])   //january 20
	//fmt.Println("december", chicken["december"]) //december 0

	//inisialisasi map
	//var chicken = map[string]int
	//chicken["one"] = 1
	//error

	//var chicken = map[string]int{}
	//chicken["one"] = 1
	//tidak error

	//chicken := map[string]int{
	//	"one": 1,
	//	"two": 2,
	//}
	//
	//fmt.Println(chicken["two"])

	//iterasi item menggunakan foreach
	//chicken := map[string]int{
	//	"january":  14,
	//	"february": 26,
	//	"march":    35,
	//	"april":    50,
	//}
	//
	//for key, val := range chicken {
	//	println(key, " \t", val)
	//}

	//Menghapus Item map
	//chicken := map[string]int{
	//	"january":  14,
	//	"february": 26,
	//	"march":    35,
	//	"april":    50,
	//}
	//
	//fmt.Println("Before delete")
	//for key, val := range chicken {
	//	fmt.Println(key, " \t:", val)
	//}
	//
	//delete(chicken, "march")
	//
	//fmt.Println("After delete")
	//for key, val := range chicken {
	//	fmt.Println(key, " \t:", val)
	//}

	//kombinasi slice dan map
	student := []map[string]string{
		map[string]string{"name": "Riko", "Gender": "Male"},
		map[string]string{"name": "Santos", "Gender": "Male"},
		map[string]string{"name": "Sodimmm", "Gender": "Male"},
	}

	for _, student := range student {
		fmt.Println("Name:", student["name"], "\n", "Gender:", student["Gender"])
	}
}
