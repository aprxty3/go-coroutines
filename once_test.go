package go_coroutines

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

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			defer group.Done()
		}()
	}

	group.Wait()
	fmt.Println("counter:", counter)
}
