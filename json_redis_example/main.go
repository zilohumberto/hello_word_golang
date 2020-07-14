package main

import (
	"github.com/zilohumberto/hello_word_golang/json_redis_example/pkg/json"
	"log"
)

func main(){
	service := json.NewService("ws-core-messages")
	messageOne := json.Message{
		Action: "hi-world",
		Body: "raw message",
	}
	err := service.SetMessage(&messageOne)
	if err != nil{
		log.Panic(err)
	}
	message, err := service.GetMessage()
	if err != nil{
		log.Panic(err)
	}
	log.Printf("A message was found\n%s %s\n", message.Action, message.Body)
}
