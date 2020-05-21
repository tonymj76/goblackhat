package main

import (
	"fmt"
	"net"
	"sort"
)

func walker(streamPort <-chan int, streamResult chan<- int) {
	for p := range streamPort {
		c, err := net.Dial("tcp", fmt.Sprintf("scanme.nmap.org:%d", p))
		if err != nil {
			streamResult <- 0
			continue
		}
		c.Close() //only when there is connection you can close the connection
		streamResult <- p
	}
}

func main() {
	posts := []int{}
	n := 1024
	streamPost := make(chan int, 100)
	streamResult := make(chan int)

	for i:=0; i < cap(streamPost); i++ {
		go walker(streamPost, streamResult)
	}
	
	go func() {
		for i:=1; i< n; i++{
			streamPost <- i
		}
	}()

	for i:=1; i< n; i++{
		post := <-streamResult
		if post != 0 {
			posts = append(posts, post)
		}
	}
	close(streamPost)
	close(streamResult)
	fmt.Println(sort.Ints(posts))
}
