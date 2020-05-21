package main

import (
	"fmt"
	"net"
)

func simpleScan() error {
	_, err := net.Dial("tcp", "scanme.nmap.org:22")
	if err != nil {
		return err
	}
	return nil
}
func simpleScan2() {
	for i := 1; i < 1024; i++ {
		_, err := net.Dial("tcp", fmt.Sprintf("scanme.nmap.org:%d", i))
		fmt.Println(i)
		if err == nil {
			fmt.Println("we have a connection")
			return
		}
	}
}

func main() {
	err := simpleScan()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("we have a connection")
	simpleScan2()
}
