package main


import (
	"fmt"
	"net/http"
)

func connect(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Connected successfully"))
}

func main(){
	http.HandleFunc("/connect", connect)
	err:= http.ListenAndServeTLS(":4433", "./certs/server.crt", "./certs/server.key", nil)
	if err!=nil{
		fmt.Println("Failed to connect", err)
	}
}
