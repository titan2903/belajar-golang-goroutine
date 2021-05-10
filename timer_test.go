package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	//! secara manual
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())
	time := <- timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	//! lebih simple
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())
	time := <- channel
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}

	group.Add(1)

	time.AfterFunc(5 * time.Second, func() {
		fmt.Println("in after func",time.Now())
		group.Done()
	})
	fmt.Println(time.Now())
	group.Wait()
}