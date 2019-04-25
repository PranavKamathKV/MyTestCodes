package main

import(
	"fmt"
	"net/http"
	"time"
)

type timeHandler struct{
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	tm:= time.Now().Format(th.format)
	w.Write([]byte("The time is : "+ tm))
}

func main(){
	mux := http.NewServeMux()

	th:=&timeHandler{format: time.RFC1123}
	mux.Handle("/time",th)

	th3:= &timeHandler{format: time.RFC3339}
	mux.Handle("/time/rfc3339",th3)

	fmt.Println("Listening to 5500 as usual")
	http.ListenAndServe(":5500",mux)


}
