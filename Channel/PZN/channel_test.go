package pzn_test

import (
	"fmt"
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