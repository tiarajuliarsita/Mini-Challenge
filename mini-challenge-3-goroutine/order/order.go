package main

import (
	"fmt"
	"sync"
)

func main() {
	bisa := []string{"bisa1", "bisa2", "bisa3"}
	coba := []string{"coba1", "coba2", "coba3"}

	c := make(chan interface{}, 4)
	var wg sync.WaitGroup
	var mt sync.Mutex
	
	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go Print(bisa, c, &wg, &mt)
		c <- i
		go Print(coba, c, &wg, &mt)
		c <- i
	}
	wg.Wait()
}

func Print(s []string, c chan interface{}, wg *sync.WaitGroup, mt *sync.Mutex) {
	// var mutex sync.Mutex
	defer wg.Done()
	mt.Lock()
	fmt.Println(s, <-c)
	mt.Unlock()
}
