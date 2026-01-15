package sync

import (
	"sync"
	"fmt"
)

func MutexExample() {
	count := 0
	var mu sync.Mutex
	
	for i := 0; i < 5; i++ {
		go func() {
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	
	fmt.Println("Final Count:", count)
}