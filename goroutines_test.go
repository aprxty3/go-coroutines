package go_coroutines

import (
	"fmt"
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
