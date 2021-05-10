package belajargolanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)


func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	// channel <- "titan" //! Mengirim data ke channel

	// data := <- channel //! menerima data dari channel
	// fmt.Println(<- channel)


	defer close(channel) //!close channel

	go func ()  {
		time.Sleep(2 * time.Second)
		channel <- "Titanio Yudista" //! Ngeblock channel karena data di channel tidak ada yang mengambil data di channel
		fmt.Println("selesai mengirim data ke channel")
	}()

	data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

/*
	pastikan channel harus ada yang mengrim dan menerima
	jika tidak ada yang mengirim channel maka akan terjadi error deadlock
	tidak membutuhkan return value
	komunikasi antar goroutine menggunakan channel di golang
*/

func GiveMeResponse(channel chan string) { //! Channel sebagai parameter dan tidak butuh pointer
	time.Sleep(2 * time.Second)
	channel <- "Titanio Yudista"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel) //!close channel

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) { //! mengirim data  ke channel
	time.Sleep(2 * time.Second)
	channel <- "Titanio Yudista"
}

func OnlyOut(channel <-chan string) { //! mengambil data dari channel
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) //!close channel

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedCahnnel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Titanio"
	channel <- "yudi"
	channel <- "eko"

	go func ()  {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)


	go func ()  {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		defer close(channel) //! channel harus di close, jika tidak terkena deadlock
	}()

	for data := range channel {
		fmt.Println("menerima data", data)
	}

	fmt.Println("Selesai")
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
		case data:= <-channel1:
			fmt.Println("Data dari channel1", data)
			counter++
		case data:= <-channel2:
			fmt.Println("Data dari channel2", data)
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
		case data:= <-channel1:
			fmt.Println("Data dari channel1", data)
			counter++
		case data:= <-channel2:
			fmt.Println("Data dari channel2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}