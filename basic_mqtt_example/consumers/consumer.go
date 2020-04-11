package consumers

import (
	"fmt"

	"github.com/zilohumberto/hello_word_golang/basic_mqtt_example/mqtttools"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Listen(c <-chan bool) {
	uri, topic := mqtttools.GetURLTopic()
	client := mqtttools.Connect("sub", uri)
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
