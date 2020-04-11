package producers

import (
	"basic_mqtt_example/mqtt_tools"
	"time"
)

func Produce(c chan<- bool) {
	uri, topic := mqtt_tools.GetURLTopic()
	commands := [4]string{
		"client_connected",
		"writting_message",
		"message:hola mundo",
		"client_leave",
	}
	client := mqtt_tools.Connect("pub", uri)
	for c := range commands {
		time.Sleep(time.Second * 1)
		client.Publish(topic, 0, false, commands[c])
	}
	time.Sleep(time.Second * 1)
	close(c)
}
