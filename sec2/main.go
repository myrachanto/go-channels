package main

import (
	"fmt"
	"math/rand"
	"strings"
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
	producer(3,d)
	consumer(1,d)
	consumer(2,d)
	wgProducers.Wait()
	close(d)
	wgConsumers.Wait()

}
func consumer(id int, in chan string){
	wgConsumers.Add(1)
	go func(){
		db := make(map[string]int)
		var fields []string
		for v := range in {
			fields = strings.Split(v, ",")
			//the ++ is incrementing the int part of the map
			db[fields[0]]++
		}
		for k,v := range db {
			fmt.Printf("Consumer %v -proccesd %v items for %v\n",id, k,v)
		}
		wgConsumers.Done()
	}()
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