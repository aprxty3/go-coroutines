package go_coroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutines(t *testing.T) {
	go HelloWorld()
	fmt.Println("Aww")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(num int) {
	fmt.Println("Dipslay Number:", num)
}

func TestManyGoRoutines(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}

func TestChannelTest(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel <- "Hello World"
		fmt.Println("Sending data to channel")
	}()

	data := <-channel
	fmt.Println("Data from channel:", data)
	close(channel)
}

func TestChannelAsParam(t *testing.T) {
	channel := make(chan string)

	go GiveResponse(channel)
	fmt.Println("Waiting for response")
	data := <-channel
	fmt.Println("Data from channel:", data)
	close(channel)

}

func GiveResponse(channle chan string) {
	time.Sleep(1 * time.Second)
	channle <- "Hello World"
}

func OnlyIn(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "Hello World"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println("Data from channel:", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
	close(channel)
}

func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 4)
	defer close(channel)

	go func() {
		channel <- "Hello World"
		channel <- "Aji"
		channel <- "Pras"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(1 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Data from channel:", data)
	}
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveResponse(channel1)
	go GiveResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2:", data)
			counter++
		default:
			fmt.Println("No data received")
		}
		if counter == 2 {
			fmt.Println("All data received")
			return
		}
	}
}
