package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type TestTitle struct{
	Title string
	Paragraphs string
}

func testHandler(w http.ResponseWriter, r *http.Request){
	p:= TestTitle{Title: "Wassup!!", Paragraphs:"How's your day going?"}
	t, err := template.ParseFiles("basic_template.html")
	if err != nil{
		fmt.Sprintf("Oops!!! Error occurred")
		panic(err)
	}
	err= t.Execute(w,p)
	fmt.Println(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, "<h1> Its working like anything </h1")
}

func main(){
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/test", testHandler)
	http.ListenAndServe(":8000", nil)
}


