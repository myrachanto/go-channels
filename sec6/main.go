package main

import "fmt"

func main (){
	//channel selection
	fmt.Print("Random bit stream: ")
	const total = 1000
	bits := randombitsGen(total)
	m := make(map[int8]int)
	for v := range bits {
		m[v]++
	}
	for k,v := range m {
		f:= (float32(v)/total) * 100
		fmt.Printf("%v occured %.2f%% of the time\n", k, f)
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