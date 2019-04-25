package main

import(
	"fmt"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request){
	tm:=time.Now().Format(time.RFC1123)
	w.Write([]byte("Time is : "+tm))
}


func main(){
	mux := http.NewServeMux()

	th:= http.HandlerFunc(timeHandler)
	mux.Handle("/time",th)

	fmt.Println("Listening to the port 5500")
	http.ListenAndServe(":5500",mux)
}
