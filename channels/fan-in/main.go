package main

import (
	"fmt"
	"sync"
	"time"
)

func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(channels))

	for _, c := range channels {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- i
			time.Sleep(time.Millisecond * 200)
		}
		close(ch1)
	}()

	go func() {
		for i := 6; i <= 10; i++ {
			ch2 <- i
			time.Sleep(time.Millisecond * 300)
		}
		close(ch2)
	}()

	for el := range merge(ch1, ch2) {
		fmt.Println("el", el)
	}
}
