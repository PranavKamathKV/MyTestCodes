package main

import "fmt"
import "time"

func test(res chan bool){
	fmt.Println("Inside test method")

	time.Sleep(2)
	fmt.Println("Print after sleep")
	res <- true
}

func main(){
	fmt.Println("Creating the channel")
	res:= make(chan bool, 1)

	go test(res)

	<- res
}
