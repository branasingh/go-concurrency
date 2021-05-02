package main

import (
	"fmt"
	"time"
)

func main() {
	go count("Cat")
	go count("Dog")
}

func count(item string) {
	// Infinite incremental counter
	for i := 1; true; i++ {
		fmt.Println(i, item)
		time.Sleep(time.Millisecond * 500)
	}
}
