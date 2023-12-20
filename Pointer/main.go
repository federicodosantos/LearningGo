package main

import "fmt"

/*
Pointer adalah alamat memori. Maksudnya, variable pointer ini berisi alamat memori suatu nilai,
misal, ada nilai integer 4, maka pointer nya adalah alamat memori nilai itu, bukan nilai dari integer 4 itu
*/

//func main() {
//numberA := 4
//var numberB *int = &numberA
//
//fmt.Println("NumberA (value): ", numberA)    //4
//fmt.Println("NumberA (address): ", &numberA) //alamat memori, referencing(&) value ke pointer
//
//fmt.Println("NumberB (value): ", *numberB)  //4, deferencing(*) pointer ke value
//fmt.Println("NumberB (address): ", numberB) //alamat memori
//
//	var animalsA string = "Hyena"
//	var animalsB *string = &animalsA
//
//	fmt.Println("AnimalsA (value): ", animalsA)
//	fmt.Println("AnimalsA (address): ", &animalsA)
//	fmt.Println("AnimalsB (value): ", *animalsB)
//	fmt.Println("AnimalsA (address): ", animalsB)
//
//	fmt.Println()
//	animalsA = "Dog"
//
//	fmt.Println("AnimalsA (value): ", animalsA)
//	fmt.Println("AnimalsA (address): ", &animalsA)
//	fmt.Println("AnimalsB (value): ", *animalsB)
//	fmt.Println("AnimalsA (address): ", animalsB)
//}

// Pointer as Parameter
func main() {
	var animalsA string = "Hyena"
	fmt.Println("Before: ", animalsA)

	change(&animalsA, "Dog")
	fmt.Println("After : ", animalsA)
}

func change(original *string, value string) {
	*original = value
}
