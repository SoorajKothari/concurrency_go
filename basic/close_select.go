package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int,2) // 2 is basically the buffer of channel
	exit := make(chan struct{}) // this channel is to wait for program to end

	go func() {
		for i:=0; i<5; i++ {
			fmt.Println(time.Now(), i, "sending message")
			channel <- i //sending messages to channel
			fmt.Println(time.Now(), i, "message sent")

			time.Sleep(1 * time.Second)
		}

		fmt.Println(time.Now(), "all completed")

		close(channel) //no more messages can be sent after channel closes.
 	}()

	go func () {
		for{
			select{ //Select is used to listen to multiple channels in our case we only have one channel here. so we can use range.

			case value, open := <- channel:
				if !open {
					close(exit)
					return
				}
				fmt.Println(time.Now(), "received", value)
			//default:
			//	fmt.Println("Nothing is happening")	
			}

		}
		//
	}()

	fmt.Println(time.Now(), "waiting for everything to finish")
	<-  exit
	fmt.Println(time.Now(), "All done")
}
