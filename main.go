package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var firstName = "Riko"
	lastName := "keren"

	fmt.Println("My name is " + firstName + " " + lastName)

	//multi-variable
	var first, second, third string
	first, second, third = "Lionel", "Messi", "Yasallam"

	fmt.Println(first, second, third)

	one, two, three := "satu", "dua", "tiga"
	fmt.Println(one, two, three)

	satu, isSunday, komaDua, say := 1, true, 2.2, "Hallo bos"
	fmt.Println(satu, isSunday, komaDua, say)

	//underscore(_)
	_ = "Belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "john", "kosong"
	fmt.Println(name)

	//new
	nama := new(string)
	fmt.Println(*nama)
}
