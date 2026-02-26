package main

import (
	"fmt"
	"time"
)

func demoLifecycle() {
	fmt.Println("1.Start")
	go func() {
		fmt.Println("4. не виконуэться")
	}()
	fmt.Println("2.готуэться")
	fmt.Println("3. Основний поток")
	fmt.Println("5. End")

}
func main() {

	demoLifecycle()
	time.Sleep(time.Second)
}
