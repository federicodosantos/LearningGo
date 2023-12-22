package main

import (
	"errors"
	"fmt"
	"strings"
)

//PEMANFAATAN ERROR
//Error merupakah sebuah tipe

//func main() {
//	var input string
//	fmt.Print("Type some number: ")
//	fmt.Scanln(&input)
//
//	var number int
//	var err error
//	number, err = strconv.Atoi(input)
//
//	if err == nil {
//		fmt.Println(number, "is number")
//	} else {
//		fmt.Println(input, "is not number")
//		fmt.Println(err.Error())
//	}
//}

// MEMBUAT CUSTOM ERROR
//func validate(input string) (bool, error) {
//	if strings.TrimSpace(input) == "" {
//		return false, errors.New("Cannot be empty!")
//	}
//	return true, nil
//}
//
//func main() {
//	var name string
//	fmt.Print("Type your name: ")
//	fmt.Scanln(&name)
//
//	if valid, err := validate(name); valid {
//		println("Hallo", name)
//	} else {
//		fmt.Println(err.Error())
//	}
//}

// PENGGUNAAN PANIC
// Panic digunakan untuk menampilkan stack trace error sekaligus menghentikan flow goroutine
//func validate(input string) (bool, error) {
//	if strings.TrimSpace(input) == "" {
//		return false, errors.New("Cannot be empty!")
//	}
//	return true, nil
//}
//
//func main() {
//	var name string
//	fmt.Print("Type your name :")
//	fmt.Scanln(&name)
//
//	if valid, err := validate(name); valid {
//		fmt.Println("Hallo", name)
//	} else {
//		panic(err.Error())
//		fmt.Println("end")
//	}
//}

// PENGGUNAAN RECOVER
// Recover berguna untuk meng-handle panic error.
func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Cannot be empty!")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func main() {
	defer catch()

	var name string
	fmt.Print("Type your name :")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Hallo", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}
