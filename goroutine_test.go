package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)


func RunhelloWorld() {
	fmt.Println("hello world")
}

func TestCreateGoroutine(t *testing.T) {
	go RunhelloWorld() //! penggunaan goroutine
	fmt.Println("Ups...")

	time.Sleep(1 * time.Second)
}

/*
	Proses asyncronus menggunakan goroutine
	tidak cocok menggunakan goroutine jika function tersebbut mengembalikan value, karena return valuenya tidak bisa di tangkap

*/

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		DisplayNumber(i)
	}
	fmt.Println("Display")

	time.Sleep(5 * time.Second)
}