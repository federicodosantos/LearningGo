package main

import "fmt"

/*
Method adalah fungsi yang menempel pada type/struct. Method diakses melalui variable objek
*/
//
//type student struct {
//	name  string
//	grade int
//}
//
//func (s student) sayHello() {
//	fmt.Println("Hello ", s.name)
//}
//
//func (s student) getNameAt(i int) string {
//	return strings.Split(s.name, " ")[i-1]
//}
//
//func main() {
//	var s1 = student{name: "Riko ganteng", grade: 2}
//
//	s1.sayHello()
//
//	var name = s1.getNameAt(2)
//	fmt.Println("Nickname : ", name)
//}

// METHOD POINTER : method yang variable objek nya pemilik method tersebut
type student struct {
	name  string
	grade int
}

func (s student) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func main() {
	var s1 = student{name: "Riko keren", grade: 2}
	fmt.Println("Before change name : ", s1.name)

	s1.changeName1("Jamal keren")
	fmt.Println("Name : ", s1.name) //Riko keren, nilainya berubah hanya didalam method saja, tapi nilai reference objeknya tidak

	s1.changeName2("Jamil keren")
	fmt.Println("Name : ", s1.name) // Jamil keren
}
