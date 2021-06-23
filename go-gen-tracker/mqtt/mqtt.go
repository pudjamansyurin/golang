package mqtt

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ClientConfig struct {
	Host     string
	Port     int
	ClientId string
	Username string
	Password string
}

type Mqtt struct {
	client mqtt.Client
	Config ClientConfig
}

func (mq *Mqtt) Connect() error {
	opts := createClientOptions(mq.Config)
	mq.client = mqtt.NewClient(opts)

	token := mq.client.Connect()
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (mq *Mqtt) Disconnect() {
	mq.client.Disconnect(100)
}

type Subscribers map[string]mqtt.MessageHandler

func (mq *Mqtt) Subscribe(subscribers Subscribers) error {
	for topic, handler := range subscribers {
		token := mq.client.Subscribe(topic, 1, handler)

		if token.Wait() && token.Error() != nil {
			log.Printf("Subscribed to topic %s\n", topic)
			return token.Error()
		}
	}
	return nil
}
