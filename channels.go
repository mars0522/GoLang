package main

import (
	"fmt"

	"golang.org/x/text/message"
)

func main(){
	messages := make(chan string)
	go fun(){ messages <- "ping"}()

	msg := messages
	fmt.Println(msg)
}