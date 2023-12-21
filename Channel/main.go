package main

import (
	"fmt"
	"runtime"
)

/*
Channel digunakan untuk menghubungkan goroutine satu dengan goroutine lain.
Mekanisme yang dilakukan adalah serah-terima data lewat channel tersebut
Pengiriman dan penerimaan data pada channel bersifat blocking atau synchronous.
"channel <- variable", berarti sedang terjadi pengiriman data dari variabel yang berada di
kanan lewat channel yang berada di kiri
"<- channel", berarti proses penerimaan data dari channel yang di kanan, untuk disimpan ke variabel yang di kiri.
*/

//func main() {
//	runtime.GOMAXPROCS(2)
//
//	var messages = make(chan string)
//
//	var sayHelloTo = func(who string) {
//		var data = fmt.Sprintf("Hello %s", who)
//		messages <- data //sedang terjadi pengiriman dari variabel "data" melewati channel "messages
//	}
//
//	go sayHelloTo("Riko")
//	go sayHelloTo("Jamal")
//	go sayHelloTo("Jamil")
//
//	var message1 = <-messages //penerimaan data
//	fmt.Println(message1)
//
//	var message2 = <-messages
//	fmt.Println(message2)
//
//	var message3 = <-messages
//	fmt.Println(message3)
//
//}

// CHANNEL SEBAGAI TIPE DATA PARAMETER
func printMessage(what chan string) {
	fmt.Println(<-what)
}

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	for _, each := range []string{"Riko", "Jamal", "Jamil"} {
		go func(who string) {
			var data = fmt.Sprintf("Hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages) //passing by reference, yang ditranferkan yaitu pointernya, bukan datanya
	}
}
