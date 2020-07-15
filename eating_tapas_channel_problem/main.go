package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	MinMorselTime = 30 / 10
	MaxMorselTime = (3 * 60) / 10
	MinMorsel     = 5
	MaxMorsel     = 10
)
var (
	friends = []string{"Nemer", "Luis", "Huber", "Lapian", "Victor"}
	dishes = []string{"Arepa", "Tacos", "Pollo", "Guacamole", "Poroto"}
	dinner sync.WaitGroup
)

func takeMeal(friend string){
	log.Printf("%s eating", friend)
	for i:=0; i < len(dishes); i++{
		dish := dishes[rand.Intn(len(dishes) - 0)]
		log.Printf("%s taking dish %s", friend, dish)
		morsel := rand.Intn(MaxMorsel-MinMorsel) + MinMorsel
		for m:=0; m < morsel; m++{
			log.Printf("%s taking morsel %d of dish %s", friend, m, dish)
			duration := rand.Intn(MaxMorselTime-MinMorselTime) + MinMorselTime
			time.Sleep(time.Duration(duration)* time.Second)
		}
	}
	dinner.Done()
}

func main(){
	dinner.Add(len(friends))
	for i:=0; i < len(friends); i++{
		go takeMeal(friends[i])
	}
	dinner.Wait()
}
