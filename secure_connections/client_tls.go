package main

import(
	"fmt"
	"crypto/tls"
)
func main(){
	conf:= &tls.Config{InsecureSkipVerify: true,}

	conn, err:= tls.Dial("tcp","localhost:4433", conf)

	if err!=nil{
		fmt.Println("Error while connecting to the server", err)
		return
	}

	defer conn.Close()

	wr, err:= conn.Write([]byte("Hello\n"))
	if err!=nil{
		fmt.Println("Error while writing",wr,err)
		return
	}

	buf:=make([]byte, 100)
	rd, err:= conn.Read(buf)
	if err!=nil{
		fmt.Println("Error reading the response from the server", rd, err)
		return
	}

	fmt.Println(string(buf[:rd]))
}
