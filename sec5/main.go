package main

import "fmt"

func main (){
	//channel selection
	fmt.Print("Random bit stream: ")
	bits := randombitsGen(10000)
	for v := range bits {
		fmt.Print(v)
	}
	fmt.Println()
}
func randombitsGen(l int)(out chan int8){
	out = make(chan int8)
	go func(){
			for i :=0 ; i<l; i++ {
			select{
			case out <- 0:
			case out <- 1:
			}
		}
		close(out)
	}()
	return
}