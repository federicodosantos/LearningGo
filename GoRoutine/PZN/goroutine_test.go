package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func HalloBoskuh() {
	fmt.Println("Hallo boskuhh semuanyaa....")
}

func TestGoroutine(t *testing.T) {
	go HalloBoskuh()
	fmt.Println("Sok asikk bang...")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Number : ", number)
}

func TestManyGorountine(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}
}