package main

import "fmt"


// A blocker function call
func test(some string){
	for i:=0;i<3;i++{
		fmt.Println(i, some)
	}
}
// main 
func main(){
	fmt.Println("Calling test method")
	test("Something")

	// Executing go test method as go routine

	go test("New thing")

	go func(anon string){
		fmt.Println("Inside anonymous go routing", anon)
	}("Anonymous goroutine")

        fmt.Scanln()
	fmt.Println("Finished")
}
