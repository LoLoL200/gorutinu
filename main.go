package main

import (
	"fmt"
	"time"
)

var counter int

func main() {
	for range 1000 {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("Counter:", counter)
}
