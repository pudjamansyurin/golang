package handler

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Report(client mqtt.Client, msg mqtt.Message) {
	log.Printf("VCU REPORT => %s\n", msg.Payload())
}
