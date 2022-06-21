package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(ch chan int, id int) {
	time.Sleep(time.Second)
	ch <- id * id
}

type resp struct {
	accept bool
	id     int
}

func driver(ch chan resp, id int) {
	time.Sleep(2 * time.Second * time.Duration(id))
	if id%2 == 0 {
		ch <- resp{true, id}
	} else {
		ch <- resp{true, id}
	}
}

func getDriverAcceptance(d1, d2 int) int {
	ch := make(chan resp)
	defer close(ch)

	go driver(ch, d1)
	go driver(ch, d2)

	r := <-ch
	if r.accept {
		return r.id
	}

	r = <-ch
	if r.accept {
		return r.id
	}

	return 0
}

func main() {
	start := time.Now()
	doc := []int{1, 2, 3, 4, 5, 6}

	channels := make([]chan int, len(doc))
	for i := 0; i < len(doc); i++ {
		channels[i] = make(chan int)
		defer close(channels[i])
	}

	for i, d := range doc {
		go worker(channels[i], d)
	}

	results := make([]int, len(doc))

	for i, ch := range channels {
		results[i] = <-ch
	}

	fmt.Println(results, time.Since(start))
	fmt.Println(getDriverAcceptance(1, 2))
}

func waitgroup() {
	var wg sync.WaitGroup

	words := []string{"foo", "bar", "baz"}

	for _, word := range words {
		// add one backlog to the group
		wg.Add(1)
		go func(word string) {
			// Block for a second
			time.Sleep(1 * time.Second)
			// remove that backlog from the group
			defer wg.Done()
			fmt.Println(word)
		}(word)
	}

	// Keep on doing things on the main goroutine
	fmt.Println(1)
	time.Sleep(1 * time.Second)
	fmt.Println(2)
	time.Sleep(1 * time.Second)
	fmt.Println(3)
	// This waits (block) for all goroutines to finish.
	wg.Wait()
}

func channel() {
	words := []string{"foo", "bar", "baz"}
	// Create a channel for communication
	done := make(chan bool)
	// It is nice to close the channels, like files,
	// after it's used or you may get leaks
	defer close(done)
	for _, word := range words {
		go func(url string) {
			// block for a sec
			time.Sleep(1 * time.Second)
			fmt.Println(word)

			// send signal to the channel
			done <- true
		}(word)
	}
	// Do what you have to do
	fmt.Println(1)
	time.Sleep(1 * time.Second)
	fmt.Println(2)
	time.Sleep(1 * time.Second)
	fmt.Println(3)
	// This blocks and waits
	<-done
}
