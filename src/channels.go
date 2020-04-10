package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	go send(c)
	wg.Add(1)
	go read(c)
	wg.Wait()
	fmt.Println("finalizando!")

}
func read(c <-chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	wg.Done()
}

func send(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i + 1
		time.Sleep(time.Second * 1)
	}
	close(c)
}
