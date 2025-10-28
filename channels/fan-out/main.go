package main

import (
	"fmt"
	"time"
)

func worker(in <-chan int, out chan<- int) {
	for d := range in {
		time.Sleep(time.Second)
		out <- d * 2
	}
}

func main() {
	data := make(chan int, 10)
	out := make(chan int, 10)

	for w := 1; w <= 3; w++ {
		go worker(data, out)
	}

	for j := 1; j <= 9; j++ {
		data <- j
	}
	close(data)

	for a := 1; a <= 9; a++ {
		fmt.Println("el", <-out)
	}
}
