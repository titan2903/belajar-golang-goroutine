package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)


var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func ()  {
			group.Add(1)
			once.Do(OnlyOnce) //! Hanya boleh function yang tidak boleh memiliki parameter
			defer group.Done()	
		}()
	}

	group.Wait() //! menunggu program goroutine
	fmt.Println("counter:", counter)
}