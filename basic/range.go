package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int,2)
	exit := make(chan struct{})

	go func() {
		for i:=0; i<5; i++ {
			fmt.Println(time.Now(), i, "sending message")
			channel <- i
			fmt.Println(time.Now(), i, "message sent")

			time.Sleep(1 * time.Second)
		}

		fmt.Println(time.Now(), "all completed")

		close(channel)
 	}()

	go func ()  {
		for value:= range channel { //Range and select works exactly the same way range is just used with one channel and select is used with multiple channels
			fmt.Println(time.Now(), "value", value)
		}
		close(exit)
	}()

	fmt.Println(time.Now(), "waiting for everything to finish")
	<-  exit
	fmt.Println(time.Now(), "All done")
}
