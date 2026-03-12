package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: Work cencelled : %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("%s: Working ...\n", name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go doWork(ctx, "Worker 1")
	go doWork(ctx, "Worker 2")

	time.Sleep(2 * time.Second)
	fmt.Println("Cancelling work ...")
	cancel()

	time.Sleep(time.Second)
}
