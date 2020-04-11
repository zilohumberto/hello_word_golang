package main

import (
	"basic_mqtt_example/consumers"
	"basic_mqtt_example/producers"
)

func main() {
	//it could be two diff programs
	c := make(chan bool)
	go producers.Produce(c)
	consumers.Listen(c)
}
