package go_coroutines

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func RunAsyncrho(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	fmt.Println("RunAsyncrho")

	time.Sleep(1 * time.Second)
}

func TestRunAsynchro(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncrho(group)
	}

	group.Done()
	log.Println("Done")
}
