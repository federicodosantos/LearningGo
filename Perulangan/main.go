package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	//for tanpa argumen
	fmt.Println("")
	var i = 0

	for {
		fmt.Println("Angka", i)
		i++
		if i == 5 {
			break
		}
	}

	//for range
	fmt.Println("")
	var xs = "123" //string
	for i, v := range xs {
		fmt.Println("index=", i, "value=", v)
	}
	fmt.Println()

	var ys = [5]int{10, 20, 30, 40, 50} //array
	for _, v := range ys {
		fmt.Println("Value=", v)
	}
	fmt.Println()

	var zs = ys[0:2]
	for _, v := range zs {
		fmt.Println("Value=", v)
	}
	fmt.Println()

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2}
	for k, v := range kvs {
		fmt.Println("Key=", k, "Value=", v)
	}
	fmt.Println()
}
