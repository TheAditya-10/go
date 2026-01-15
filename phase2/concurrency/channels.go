package concurrency

import "fmt"

func ChannelExample() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from the goroutine!"
	}()

	msg := <-ch
	fmt.Println(msg)
}