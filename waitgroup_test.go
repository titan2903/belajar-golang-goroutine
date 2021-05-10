package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)


func RunAsync(group *sync.WaitGroup) {
	defer group.Done() //! Selalu memanggil proses Done agar tidak terjadi deadlock
	
	group.Add(1) //! Menambahkan satu goroutine

	fmt.Println("Hallo")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i:= 0; i < 100; i++ {
		go RunAsync(group)
	}

	group.Wait() //! menunggu sampai program selesai semua di eksekusi
	fmt.Println("selesai")
}