package main

import "fmt"

func main(){
	fmt.Println("Creating the channel")

	msg1:= make(chan string)
	go func(){
		msg1 <- "What"
	}()

	msg2:= <- msg1
	fmt.Println(msg2)
}
