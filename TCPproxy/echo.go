package main

import (
	"errors"
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	//creating a buffa
	buf := make([]byte, 512)
	for {
		size, err := conn.Read(buf[:])
		if err != nil {
			log.Println("some error occured ", err.Error())
			break
		}

		if errors.Is(err, io.EOF) {
			log.Println("connection closed")
			break
		}

		log.Printf("Received %d bytes: %s", size, string(buf))
		//Sending data
		if _, err := conn.Write(buf[:size]); err != nil {
			log.Println("failed to write to conn")
			break
		}
	}
}

func main() {
	port := ":20099"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("cant Listen in port ", port)
	}
	log.Println("Listening on port... ", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("unable to accept connection")
		}
		log.Println("connection made")
		go echo(conn)
	}
}
