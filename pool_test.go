package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)


func TestPool(t *testing.T) {
	// pool := sync.Pool{} //! tanpa default value
	pool := sync.Pool{ //! menggunakan default value
		New: func () interface{}  {
			return "New"
		},
	}

	
	pool.Put("eko")
	pool.Put("kurniawan")
	pool.Put("khannedy")

	for i := 0; i < 10; i++ {
		go func ()  {
			data :=	pool.Get() //! mengambil data dalam pool
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data) //! setelah data di gunakan maka harus di kembalikan lagi datanya ke dalam pool nya, jika tidak otomatis data akan hilang dari poolnya
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("selesai")
}