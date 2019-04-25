package main

import(
	"fmt"
	"net/http"
	"time"
)

func timeHandler(format string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		tm:=time.Now().Format(format)
		w.Write([]byte("The time is : "+ tm))
	}
}

func main(){
	mux:= http.NewServeMux()
	mux.HandleFunc("/time", timeHandler(time.RFC1123))
	fmt.Println("Listening to 5500")
	http.ListenAndServe(":5500", mux)
}
