package go_coroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestGetGomaxprocs(t *testing.T) {
	group := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			fmt.Println("Hello World")
			group.Done()
		}()

	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	totalThreads := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Threads", totalThreads)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines", totalGoroutines)

	group.Wait()
}

func TestChangeThreads(t *testing.T) {
	group := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			fmt.Println("Hello World")
			group.Done()
		}()

	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThreads := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Threads", totalThreads)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines", totalGoroutines)

	group.Wait()
}
