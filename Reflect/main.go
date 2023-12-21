package main

import (
	"fmt"
	"reflect"
)

/*
Reflect adalah teknik untuk inspeksi variable, mengambil informasi dari variable,
atau bahkan memanipulasinya
*/

//func main() {
//	var number = 23
//	var reflectValue = reflect.ValueOf(number)
//	var reflectType = reflect.TypeOf(number)
//
//	if reflectValue.Kind() == reflect.Int {
//		fmt.Println("Value : ", reflectValue)
//	}
//	fmt.Println("Type : ", reflectType)
//}

// PENGAKSESAN INFORMASI PROPERTY VARIABEL OBJEK
// reflect bisa mengambil semua informasi semua property variable objek

//type student struct {
//	Name  string
//	Grade int
//}
//
//func (s *student) getPropertyInfo() {
//	var reflectValue = reflect.ValueOf(s)
//
//	if reflectValue.Kind() == reflect.Ptr {
//		reflectValue = reflectValue.Elem()
//	}
//
//	var reflectType = reflectValue.Type()
//
//	for i := 0; i < reflectValue.NumField(); i++ {
//		fmt.Println("Nama : ", reflectType.Field(i).Name)
//		fmt.Println("Tipe data : ", reflectType.Field(i).Type)
//		fmt.Println("Nilai : ", reflectValue.Field(i).Interface())
//		fmt.Println("")
//	}
//}
//
//func main() {
//	var s1 = &student{
//		Name:  "Riko",
//		Grade: 3,
//	}
//	s1.getPropertyInfo()
//}

// Pengaksesan Informasi Method Variable Objek
type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var s1 = &student{
		Name:  "Riko",
		Grade: 2,
	}
	fmt.Println("Nama : ", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Jamal"),
	})

	fmt.Println("Nama : ", s1.Name)
}
