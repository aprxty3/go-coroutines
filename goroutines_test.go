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
