// take it from: https://www.golangprograms.com/go-language/concurrency.html
package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"sync"
	"time"
)

const (
	hunger = 3
	think = time.Second / 100
	eat = time.Second / 100
)
var (
	ph = []string{"Mark", "Russell", "Rockey", "Haris", "Root" }

)
var dining sync.WaitGroup

func diningProblem(phName string, domainHand, otherHand *sync.Mutex){
	fmt.Println(phName, "Seated")
	h :=fnv.New64a()
	h.Write([]byte(phName))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration){
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	for h := hunger; h > 0; h--{
		fmt.Println(phName, "hungry")
		domainHand.Lock()
		otherHand.Lock()
		fmt.Println(phName, "Eating")
		rSleep(eat)
		domainHand.Unlock()
		otherHand.Unlock()
		fmt.Println(phName, "Thinking")
		rSleep(think)
	}
	fmt.Println(phName, "Satisfied")
	dining.Done()
	fmt.Println(phName, "Left the table")
}

func main(){
	fmt.Println("Table empty")
	dining.Add(5)
	fork0 := &sync.Mutex{}
	forkLeft := fork0

	for i:=1; i<len(ph); i++{
		forkRight := &sync.Mutex{}
		go diningProblem(ph[i], forkLeft, forkRight)
		forkLeft = forkRight
	}
	go diningProblem(ph[0], fork0, forkLeft)
	dining.Wait()
	fmt.Println("Table empty")
}
