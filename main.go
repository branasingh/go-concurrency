package main

import (
	"fmt"
	"strconv"
)

func main() {

	jobs := make(chan int, 100)
	results := make(chan string, 100)

	go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

// Jobs Channel is only can read from
// Results Channel is only can write in to
func worker(jobs <-chan int, results chan<- string) {
	for n := range jobs {
		results <- "Square root of " + strconv.Itoa(n) + " is :" + strconv.Itoa(square(n))
	}
}

func square(number int) int {
	// if number%4 == 0 {
	// time.Sleep(time.Millisecond * 5000)
	// }
	return number * number
}
