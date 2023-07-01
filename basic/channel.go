package main

import (
	"fmt"
	"time"
)

// channel will block execution until receives the value

func main() {
	channel := make(chan int)

	go func ()  {
		fmt.Println(time.Now(), "sleeping!")
		time.Sleep(2 * time.Second)

		channel <- 1000
	}()

	fmt.Println(time.Now(), "waiting for value")

	value := <- channel

	fmt.Println(time.Now(), "value is ", value)
}
