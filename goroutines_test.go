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
