package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":3400")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		data, err2 := ioutil.ReadAll(bufio.NewReader(conn))
		if err2 != nil {
			panic(err2)
		}
		fmt.Println("line1", string(data))
		conn.Close()
		//time.Sleep(time.Second)
	}
}
