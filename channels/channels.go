package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	go func() { time.Sleep(30); messages <- "ping" }()
	go func() { messages <- "ping twice" }()
	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
}
