package go_coroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()

	fmt.Println("WaitCondition:", value)
	cond.L.Unlock()

}
func TestCondition(t *testing.T) {

	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 + time.Second)
			cond.Signal()
		}
	}()

	//go func() {
	//	group.Add(1)
	//	time.Sleep(1 + time.Second)
	//	cond.Broadcast()
	//}()

	group.Wait()
}
