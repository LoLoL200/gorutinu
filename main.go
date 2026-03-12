package main

import (
	"context"
	"fmt"
	"time"
)

func taskWithDeadline(ctx context.Context, taskName string) {
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Printf("%s: Dead see for %v\n", taskName, deadline.Format("15:04:05"))
	}

	select {
	case <-time.After(2 * time.Second):
		fmt.Printf("%s: Task completed\n", taskName)
	case <-ctx.Done():
		fmt.Printf("%s:Task cancelled due to deadline: %v\n", taskName, ctx.Err())
	}

}

func main() {
	deadline := time.Now().Add(1 * time.Second)
	ctx, cencel := context.WithDeadline(context.Background(), deadline)
	defer cencel()

	go taskWithDeadline(ctx, "Task1")
	go taskWithDeadline(ctx, "Task2")
	time.Sleep(3 * time.Second)
}
