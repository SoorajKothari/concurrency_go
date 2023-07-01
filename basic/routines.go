package main

import "fmt"

//Since routines don't wait we will only see the print in main on running this programming

func main() {
	go routine() 
	fmt.Println("Hello after routine")

}

func routine() {
	fmt.Println("Hello in routine")
}
