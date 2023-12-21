package main

import "fmt"

/*
Struct merupakan class jika di bahasa go, tetapi ada perbedaan diantara keduanya
*/

//type student struct {
//	name string
//	age  int
//}
//
//func main() {
//	var s1 student
//
//	s1.name = "Riko"
//	s1.age = 25
//
//	fmt.Println("Nama: ", s1.name)
//	fmt.Println("Umur: ", s1.age)
//}

// INISIALISASI OBJECT STRUCT
//type student struct {
//	name string
//	age  int
//}
//
//func main() {
//	var s1 = student{}
//	s1.name = "Riko"
//	s1.age = 25
//
//	var s2 = student{"Riko", 25}
//
//	var s3 = student{
//		name: "Riko",
//		age:  25,
//	}
//
//	fmt.Println("Student 1: ", s1.name)
//	fmt.Println("Student 2: ", s2.name)
//	fmt.Println("Student 3: ", s3.name)
//}

// VARIABEL OBJEK POINTER
//type student struct {
//	name string
//	age  int
//}
//
//func main() {
//	var s1 = student{"Riko", 22}
//
//	var s2 *student = &s1
//	fmt.Println("Student 1: ", s1.name)
//	fmt.Println("Student 2: ", s2.name) //Riko, inilah keistimewaan property dalam objek pointer
//
//	s2.name = "Jamal"
//	fmt.Println("Student 1: ", s1.name)
//	fmt.Println("Student 2: ", s2.name)
//}

// EMBEDED STRUCT : mekanisme untuk menempelkan sebuah struct sebagai properti struct lain.
//type person struct {
//	name string
//	age  int
//}
//
//type student struct {
//	person
//	grade int
//}
//
//func main() {
//	var s1 = student{
//		person: person{"Riko", 22},
//		grade:  4,
//	}
//
//	fmt.Println("Student name: ", s1.name)
//	fmt.Println("Student age: ", s1.age)
//	fmt.Println("Student grade: ", s1.grade)
//}

// ANONYMUS STRUCT : struct yang dideklarasikan pada saat pembuatan object
// teknik ini cukup efisien untuk pembuatan object sekali pakai saja
//type person struct {
//	name string
//	age  int
//}
//
//func main() {
//	var s1 = struct {
//		person
//		grade int
//	}{} //{}->ini property, harus ada, jika tidak bakal terjadi error
//
//	s1.person = person{name: "Riko", age: 22}
//	s1.grade = 10
//
//	fmt.Println("Student name: ", s1.name)
//	fmt.Println("Student age: ", s1.age)
//	fmt.Println("Student grade: ", s1.grade)
//}

// KOMBINASI SLICE DAN STRUCT
//type person struct {
//	name string
//	age  int
//}
//
//func main() {
//	var allStudents = []person{
//		{name: "Riko", age: 22},
//		{name: "Jamal", age: 32},
//		{name: "Jamil", age: 42},
//	}
//
//	for _, e := range allStudents {
//		fmt.Println(e.name, " age is ", e.age)
//	}
//}

// INISIALISASI SLICE ANONYMUS STRUCT
//type person struct {
//	name string
//	age  int
//}
//
//func main() {
//	var allStudents = []struct {
//		person
//		grade int
//	}{
//		{person: person{name: "Riko", age: 22}, grade: 12},
//		{person: person{name: "Jamal", age: 32}, grade: 11},
//		{person: person{name: "Jamil", age: 42}, grade: 10},
//	}
//
//	for _, student := range allStudents {
//		fmt.Println(student)
//	}
//}

// NESTED STRUCT : Teknik ini biasa digunakan ketika decoding data json yang struktur datanya cukup
//kompleks dengan proses decode hanya sekali.
//type student struct {
//	person struct {
//		name string
//		age  int
//	}
//	grade int
//}

// TYPE ALIAS
type person struct {
	name string
	age  int
}
type people = person

func main() {
	var p1 = people{
		name: "Riko",
		age:  22,
	}

	fmt.Println(p1.name)
	fmt.Println(p1.age)
}
