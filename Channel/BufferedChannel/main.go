package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
Pada umumnya proses transfer data lewat channel dilakukan dengan cara unbuffered.
Agar pengiriman data bisa dilakukan secara asynchronous, maka dibuatlah buffered channel.
Buffered channel adalah channel yang menentukan angka buffernya
atau bisa dibilang angka penentu jumlah data yang bisa dikirimkan secara bersamaan.
Akan tetapi, penerimaan data SELALU dilakukan secara synchronous
*/

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 3) //jika 3, berarti jumlah maksimal buffer 4. (input + 1)

	go func() {
		for {
			i := <-messages
			fmt.Println("Receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Send data", i)
		messages <- i
	}

	time.Sleep(1 * time.Second) // menambahkan delay sebelum statement berikutnya dieksekusi
}
