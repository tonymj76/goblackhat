package main

import (
	"bufio"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Println("some error occured ", err.Error())
	}
	log.Printf("Read %d bytes: %s\n", len(s), s)
	log.Println("Writing data")
	wr := bufio.NewWriter(conn)
	b, err := wr.WriteString(s)
	if err != nil {
		log.Fatalln("Can't write")
	}
	log.Printf("Write %d bytes\n", b)
	wr.Flush()

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
