package main

import (
	"design-patterns-Golang/Singleton/singleton"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			// done with worker
			defer wg.Done()
			s := singleton.GetInstance(i)
			fmt.Printf("%d: Data: %s\n", s.Info, s.Data)
		}()
	}

	// wait until all workers done their job
	wg.Wait()
}
