package concurrency

import (
	"fmt"
	"time"
)

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(150 * time.Millisecond)
	}
}