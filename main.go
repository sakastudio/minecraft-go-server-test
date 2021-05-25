package main

import (
	"bufio"
	"fmt"
	"net"
)

type response struct {
	payload string
	id int
}

func main() {
	ln,err := net.Listen("tcp",":25565")
	if err != nil{
		panic(err)
	}
	for {
		conn,err := ln.Accept()
		if err!= nil{

		}
		go func(){
			fmt.Println(conn.RemoteAddr())
			fmt.Println(bufio.NewReader(conn))

			write, err := conn.Write([]byte("newmessage" + "\n"))
			if err != nil {
				return
			}
			fmt.Println(write)
		}()
	}
}
