package main

import(
	"fmt"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request){
	tm:= time.Now().Format(time.RFC1123)
	w.Write([]byte("Time is : " + tm))
}


func main(){
	mux:= http.NewServeMux()

	mux.HandleFunc("/time", timeHandler)

	fmt.Println("Listening to 5500")
	http.ListenAndServe(":5500",mux)
}
