package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func send(words []string, ch chan string) {
	for _, word := range words {
		conn, err := net.Dial("tcp", "127.0.0.1:3400")
		if err != nil {
			panic(err)
		}
		conn.Write(([]byte)(word))
		time.Sleep(time.Second)
		fmt.Println("Sent", word)
		conn.Close()
	}
	ch <- "done"
}

func parseCommandLine(ch chan []string) {
	words := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word := scanner.Text()
		if word == "Quit" {
			break
		}
		words = append(words, word)
	}
	ch <- words
}

func main() {

	chan1 := make(chan []string)
	chan2 := make(chan string)
	go parseCommandLine(chan1)
	words := <-chan1
	go send(words, chan2)
	<-chan2
}
