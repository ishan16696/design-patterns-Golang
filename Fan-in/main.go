package main

import (
	"fmt"
	"sync"
	"time"
)

func search(query, searchEngine string) <-chan string {
	result := make(chan string)
	go func() {
		defer close(result)

		if searchEngine == "Google" {
			time.Sleep(80 * time.Millisecond)
		} else {
			time.Sleep(120 * time.Millisecond)
		}
		result <- fmt.Sprintf("%s result for %q", searchEngine, query)
	}()

	return result
}

// Merge different channels in one channel
func fanIn(inputs ...<-chan string) <-chan string {
	var wg sync.WaitGroup

	out := make(chan string)
	wg.Add(len(inputs))

	// start a goroutine for each input channel.
	// send copies all values from channel c to out, then calls wg.Done.
	send := func(c <-chan string) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	for _, c := range inputs {
		go send(c)
	}

	// Start a goroutine to close channel "out" once all the send goroutines are done.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	result1 := search("SomeThing", "Google")
	result2 := search("SomeThing", "Bing")
	result3 := search("SomeThing", "Perplexity")

	combineResult := fanIn(result1, result2, result3)

	for result := range combineResult {
		fmt.Println(result)
	}
}
