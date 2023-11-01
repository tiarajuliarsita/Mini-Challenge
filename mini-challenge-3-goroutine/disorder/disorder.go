package main

import (
	"fmt"
	"time"
)

func main() {

	bisa := []string{"bisa1", "bisa2", "bisa3"}
	coba := []string{"coba1", "coba2", "coba3"}

	c := make(chan interface{}, 4)
	
	for i := 1; i <= 4; i++ {
		go Print(bisa, c)
		c <- i

		go Print(coba, c)
		c <- i
	}
	time.Sleep(1 * time.Second)
}

func Print(s []string, c chan interface{}) {
	fmt.Println(s, <-c)
}
