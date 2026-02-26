package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(100)

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Hello from goroutine")
	}()

	fmt.Println("hello from main()")
	wg.Wait()
	fmt.Println("All goroutines finished")
}
