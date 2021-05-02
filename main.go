package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string)

	go count("Cat", channel)

	for {
		msg, open := <-channel
		if !open {
			break
		}
		fmt.Println(msg)
	}

}

func count(item string, channel chan string) {
	// Infinite incremental counter
	for i := 1; i <= 5; i++ {
		channel <- item
		time.Sleep(time.Millisecond * 500)
	}
	close(channel)
}
