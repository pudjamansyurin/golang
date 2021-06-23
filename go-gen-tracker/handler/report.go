package handler

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pudjamansyurin/go-gen-tracker/packet"
	// "github.com/pudjamansyurin/go-gen-tracker/decoder"
)

func Report(client mqtt.Client, msg mqtt.Message) {
	// decoder.Report(msg.Payload())
	for _, h := range packet.Header {
		fmt.Println(h.Title)
	}
	hexString := strings.ToUpper(hex.EncodeToString(msg.Payload()))
	log.Printf("[REPORT] %s\n", hexString)
}
