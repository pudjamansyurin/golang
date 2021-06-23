package mqtt

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func createClientOptions(config ClientConfig) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.Host, config.Port))
	opts.SetClientID(config.ClientId)
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)

	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectionLostHandler

	return opts
}

func messagePubHandler(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Message %s received on topic %s\n", msg.Payload(), msg.Topic())
}

func connectHandler(client mqtt.Client) {
	log.Println("Connected")
}

func connectionLostHandler(client mqtt.Client, err error) {
	log.Printf("Connection Lost: %s\n", err.Error())
}
