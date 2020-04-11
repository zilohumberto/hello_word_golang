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
	// the same result but with select
	cc := make(chan int)
	go sendSelect(cc)
	readSelect(cc)
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

func readSelect(c <-chan int) {
	fmt.Println("start: readSelect")
	for {
		select {
		case p, ok := <-c:
			if ok {
				fmt.Println(p)
			} else {
				return
			}
		}
	}
}

func sendSelect(c chan<- int) {
	fmt.Println("start: sendSelect")
	for i := 0; i < 5; i++ {
		c <- i + 1
		time.Sleep(time.Second * 1)
	}
	close(c)
}
