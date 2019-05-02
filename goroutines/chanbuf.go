package main

import "fmt"

func main(){
	fmt.Println("Creating the channel which can recieve upto two values")
	message:= make(chan string, 2)

	fmt.Println("Sending the message into the channel")

	go func(){
		message <- "Something is"
		message <- "Better than nothing"
	}()

	fmt.Println("Printing the channel values....\n\n\n")
	fmt.Println(<-message, <-message)
}
