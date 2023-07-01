package main

import "fmt"

  /*
  Since go routine gives us concurrency hence there will be a listener before we assign value to channel
  */

func main() {
	channel := make(chan int)

	go func() {
		channel <- 123
	}()

	value := <-channel
	fmt.Println(value)
}
