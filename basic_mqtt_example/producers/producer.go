package producers

import (
	"time"

	"github.com/zilohumberto/hello_word_golang/basic_mqtt_example/mqtttools"
)

func Produce(c chan<- bool) {
	uri, topic := mqtttools.GetURLTopic()
	commands := [4]string{
		"client_connected",
		"writting_message",
		"message:hola mundo",
		"client_leave",
	}
	client := mqtttools.Connect("pub", uri)
	for c := range commands {
		time.Sleep(time.Second * 1)
		client.Publish(topic, 0, false, commands[c])
	}
	time.Sleep(time.Second * 1)
	close(c)
}
