package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	func() {
		go count("Cat")
	}()

	wg.Wait()
}

func count(item string) {
	// Infinite incremental counter
	for i := 1; i <= 5; i++ {
		fmt.Println(i, item)
		time.Sleep(time.Millisecond * 500)
	}
}
