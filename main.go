package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cencel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cencel()

	done := make(chan struct{})
	go func() {
		fmt.Println("Почав роботу...")
		time.Sleep(3 * time.Second)
		fmt.Println("Почав роботу...")
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Операция завершилась успишно")
	case <-ctx.Done():
		fmt.Println("Таймаут! Операция перервана", ctx.Err())

	}
}
