package go_coroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				atomic.AddInt64(&x, 1)
			}
			defer group.Done()
		}()
	}
	group.Wait()
	fmt.Println("counter:", x)
}
