package main

import(
	"fmt"
	"crypto/tls"
	"net"
	"bufio"
)

func connectionHandler(conn net.Conn){
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for{
		msg, err:= reader.ReadString('\n')
		if err!=nil{
			fmt.Println("Error while reading the message", err)
			return
		}

	        fmt.Println(msg)

	        cn, err:= conn.Write([]byte("world\n"))
	        if err!=nil{
		        fmt.Println(cn,err)
			return
	        }
	}
}


func main(){
	cert, err:= tls.LoadX509KeyPair("./certs/server.crt", "./certs/server.key")
	if err!=nil{
		fmt.Println("Error loading certificate file", err)
		panic(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	lstn, err := tls.Listen("tcp", ":4433", config)
	if err!=nil{
		panic(err)
	}

	defer lstn.Close()


	for{
		conn, err:= lstn.Accept()
		if err!=nil{
			fmt.Println("Connection error", err)
			continue
		}

		go connectionHandler(conn)
	}
}


