package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
)
var (
	
	wgProducers sync.WaitGroup
	wgWorkers sync.WaitGroup
	numberofworkers int
)

type result struct{
	name string
	weeks int
	days int
} 
func main(){
	c := make(chan result, 3)
	m := map[string]int { "Robert" : 30, "John" : 475, "Elly" : 1022, "Bob" : 99 }

	c = converter(m,c)
	go worker(c)
	go worker(c)
	go worker(c)
	wgProducers.Wait()
	wgWorkers.Wait()
	fmt.Println("The number of workers is : "+strconv.Itoa(numberofworkers))
}
func worker(c chan result)int{
	numberofworkers++
	wgWorkers.Add(1)

	go func(){
		for msg := range c{
			fmt.Println(msg.name + " has worked here for "+ strconv.Itoa(msg.weeks) + " Weeks and " + strconv.Itoa(msg.days)+ " days for the company")
	
		}
		wgWorkers.Done()
		}()
		// counter(numberofworkers)
		return numberofworkers
}
func converter(m map[string]int, c chan result) (chan result){
	wgProducers.Add(1)
 go func(){
	res := result{} 
		for k,v := range m {
			res.weeks = v/7
			res.name = k
			res.days = v%7

			f := res.weeks
			e := float64(f)
			t := math.Floor(e)
			y := int(t)
			res.weeks = y

	c <- res
	}
	close(c)
	wgProducers.Done()
 }()
	return c
}