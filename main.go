package main

import (
	"fmt"
	"time"
)

func main() {
	msg1 := make(chan string)
	msg2 := make(chan string)
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			msg1 <- "Прошло 0.5 секунды"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			msg2 <- "Прошло 2 секунды"
		}
	}()
	for {
		select {
		case m1 := <-msg1:
			fmt.Println(m1)
		case m2 := <-msg2:
			fmt.Println(m2)
		}
	}
}
