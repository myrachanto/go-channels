package main

import (
	"fmt"
	"time"
)


func main (){
	//channel selection
	d:= producer()
	consumer(d)
}
func consumer(in chan string){
	for {
		alarm := time.After(2  * time.Millisecond)
		select {
		case m, ok := <-in:
			if !ok {
				fmt.Println("No more data form in")
				return
			}
			fmt.Printf("Consumer got -%v", m)
		case <-alarm:
			fmt.Printf("TimedOut waiting on for data")
			return
		}
	}
}
func producer()(out chan string){
	out = make(chan string)
	go func(){
		count := 1

		for i := 0; i< 10; i++ {
			out <- fmt.Sprintf("Producer message %v\n", count)
			count++
		}
		time.Sleep(3 * time.Millisecond)
		for i := 0; i< 10; i++ {
			out <- fmt.Sprintf("Producer  message %v\n", count)
			count++
		}
		close(out)
	}()
	return
}
// func randombitsGen(l int)(out chan int8){
// 	out = make(chan int8)
// 	go func(){
// 			for i :=0 ; i<l; i++ {
// 			select{
// 			case out <- 0:
// 			case out <- 1:
// 			}
// 		}
// 		close(out)
// 	}()
// 	return
// }