package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listen, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalln(err)
	}
	defer listen.Close()
	for i := 0; i < 3; i++ {
		buf := make([]byte, 1024)
		n, err := listen.Read(buf)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(strings.ToUpper(string(buf[:n])))
	}
}
