package main

import(
	"fmt"
	"net/http"
)

func main(){
	mux:= http.NewServeMux()

	rh:= http.RedirectHandler("http://www.google.com",307)
	mux.Handle("/foo", rh)

	fmt.Println("Listening to the port 5500")
	http.ListenAndServe(":5500",mux)
}
