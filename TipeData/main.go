package main

import "fmt"

func main() {
	//data numerik non desimal
	var positiveNumber uint8 = 99
	var negativeNumber = -123456789

	fmt.Printf("Bilangan positif: %d\n", positiveNumber)
	fmt.Printf("Bilangan positif: %d\n", negativeNumber)

	//data numerik desimal
	var decimalNumber = 2.62
	fmt.Printf("Bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("Bilangan desimal: %.3f\n", decimalNumber)

	//string
	var message string = "Riko"
	fmt.Printf("message: %s\n", message)

	//string backticks
	message = `Nama saya "Riko keren".
	Salam kenal.
	Mari belajar golang`

	fmt.Println(message)
}
