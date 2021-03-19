package main

import "fmt"

func main (){
	//channel selection
	var ch1, ch2 chan string
	select {
	case <- ch1:
		fmt.Println("got data from channel1")
	case ch2 <- "hello":
		fmt.Println("sent to channel 2")
	default:
		fmt.Println("no communication")
	}
}