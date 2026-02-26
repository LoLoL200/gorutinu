package main

import (
	"fmt"
)

func main() {
	msg := make(chan string, 3)
	msg <- "Hello"
	msg <- "World"
	msg <- "Go"

	fmt.Println(<-msg)
	fmt.Println(<-msg)
	fmt.Println(<-msg)
}
