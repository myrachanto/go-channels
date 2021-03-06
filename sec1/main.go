package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
	wgProducers sync.WaitGroup
	wgConsumers sync.WaitGroup
)

func main(){
	d := make(chan string)
	producer(1,d)
	producer(2,d)
	go consumer(d)
	wgProducers.Wait()
	close(d)
	wgConsumers.Wait()

}
func consumer(in chan string){
	wgConsumers.Add(1)
	count := 0
	for v := range in {
		count++
		fmt.Printf("Consumer got: %v\n", v)
	}
	if count ==0 {
		fmt.Println("No data received")
		return
	}
	fmt.Printf("Processed %v items\n", count)
	wgConsumers.Done()
}
func producer (id int, out chan string){
	wgProducers.Add(1)
	go func(){
		i := 1
		end := time.Now().Add(1000 * time.Millisecond)
	
		for time.Now().Before(end){
			out <- fmt.Sprintf("Producer: %v, item: %v", id, i)
			i++
		}
		wgProducers.Done()
	}()
	// close(out)
}