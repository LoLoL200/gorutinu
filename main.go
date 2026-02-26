package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	go func() {
		a := 5
		ch <- a
		ch <- 2312
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
