package main

import (
	"fmt"
	"net"
)

//Tom Steal called this approch a naive way to go about port scanner

//This is 100% faster but have a race condition at no 22 and 80
//this will have a race condition if the close func is used
func concurrentScan() <-chan string {
	streamConn := make(chan string)
	for i := 1; i < 1024; i++ {
		go func(i int) {
			conn, err := net.Dial("tcp", fmt.Sprintf("scanme.nmap.org:%d", i))
			if err == nil {
				streamConn <- "we have a connection"
				// close(streamConn) we don't close like this because others will panic due to send to a close chan
				conn.Close()
			}
		}(i)
	}
	return streamConn
}

//This is fucking damn to slow is like using a normal function with loop
//it will take like 3 minute to get the work done
func concurrentScan2() <-chan string {
	streamConn := make(chan string)
	go func() {
		defer close(streamConn)
		for i := 1; i < 1024; i++ {
			_, err := net.Dial("tcp", fmt.Sprintf("scanme.nmap.org:%d", i))
			if err == nil {
				streamConn <- "we have a connection"
			}
		}
	}()
	return streamConn
}

func concurrentScan3(i int, str chan<- string) {
	c, err := net.Dial("tcp", fmt.Sprintf("scanme.nmap.org:%d", i))
	if err == nil {
		str <- fmt.Sprintf("we have a connection from scan3: %d\n", i)
		c.Close()
	}
	return
}

func main() {
	// select {
	// case str := <-concurrentScan():
	// 	fmt.Println(str)
	// case <-time.After(3 * time.Minute):
	// 	fmt.Println("Terminating program")
	// }
	streamConn := make(chan string)
	for i := 1; i < 1024; i++ {
		go concurrentScan3(i, streamConn)
	}

	select {
	case con := <-concurrentScan():
		fmt.Println(con)
	case str := <-streamConn:
		fmt.Println(str)
	}
}
