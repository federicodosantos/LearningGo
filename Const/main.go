package main

import "fmt"

func main() {
	const firstName string = "Riko"
	fmt.Print("Halo ", firstName, "!\n")

	const lastName = "Keren"
	fmt.Print("Nice to meet you ", lastName, "!\n")

	//multi konstanta
	const (
		square         = "kotak"
		isToday  bool  = false
		numeric  uint8 = 99
		floatNum       = 6.9
	)

	const (
		today string = "Senin"
		sekarang
		isToday2 = true
	)
}
