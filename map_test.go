package go_coroutines

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, group *sync.WaitGroup, value int) {

	data.Store(value, value)

	defer group.Done()
}
func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		group.Add(1)
		go AddToMap(data, group, i)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})

}
