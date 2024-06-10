package pzn_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func ()  {
		time.Sleep(2 * time.Second)
		channel <- "Halloo boskuhhh semuanyaa...."
		fmt.Println("Channel sudah dikirim ya boskuhh...")
	}()

	data := <- channel
	fmt.Println(data)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hallo boskuh semuanya...."
}

func TestChannelParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)

}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Halloo boskuh semuanya..."
}

func OnlyOut(channel <-chan string) {
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)
	
	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func ()  {
		channel <- "Haloo"
		channel <- "Boskuhhh"
		channel <- "semuanya..."
	}()

	go func ()  {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()

	time.Sleep(5 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func ()  {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)
	

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data diterima dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data diterima dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)
	

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data diterima dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data diterima dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data bosskuh...")
		}

		if counter == 2 {
			break
		}
	}
}