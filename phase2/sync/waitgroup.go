package sync

import (
	"sync"
	"fmt"
)

func WaitGroupExample() {
	var wg sync.WaitGroup
	count := 0
	var mu sync.Mutex
	
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Final Count:", count)
}