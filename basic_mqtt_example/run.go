package main

import (
	"github.com/zilohumberto/hello_word_golang/basic_mqtt_example/consumers"
	"github.com/zilohumberto/hello_word_golang/basic_mqtt_example/producers"
)

func main() {
	//it could be two diff programs
	c := make(chan bool)
	go producers.Produce(c)
	consumers.Listen(c)
}
