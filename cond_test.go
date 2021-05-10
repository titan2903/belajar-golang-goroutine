package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)


var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()

	cond.Wait() //! menanyakan terlebih dahulu setelah melakukan locking
	//! Setelah melakukan wait pastikan ada yang mengirim signal untuk waitnya
	fmt.Println("Done", value)
		
	cond.L.Unlock()
}


func TestCond (t *testing.T) {
	for i := 0; i < 10; i++{
		go WaitCondition(i)
	}

	go func ()  {
		for i := 0; i < 10; i++{
			time.Sleep(1 * time.Second)
			cond.Signal() //! mengirim signal ke await
			//! kalau signal maka dia akan satu persatu di eksekusi
		}
	}()

	// go func ()  {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast() //! langsung semua di eksekusi jika menggunakan Broadcart()
	// }()

	group.Wait()
}
