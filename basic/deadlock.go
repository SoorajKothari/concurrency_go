package main

import "fmt"

/* Deadlock
     This will create a deadlock because go code executes sequentially so to use channels we must have some code to listen to the channel before we assign the value to it.
  */

func main() {
	channel := make(chan int)

	channel <- 090078601
	value := <-channel
	fmt.Println(value)
}
