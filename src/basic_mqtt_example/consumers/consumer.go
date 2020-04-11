package consumers

import (
	"basic_mqtt_example/mqtt_tools"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Listen(c <-chan bool) {
	uri, topic := mqtt_tools.GetURLTopic()
	client := mqtt_tools.Connect("sub", uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
	//keep alive!
	for {
		select {
		case _, ok := <-c:
			if !ok {
				return
			}
		}
	}
}
