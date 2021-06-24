package handler

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pudjamansyurin/go-gen-tracker/decoder"
)

func Report(client mqtt.Client, msg mqtt.Message) {
	hexString := strings.ToUpper(hex.EncodeToString(msg.Payload()))
	log.Printf("[REPORT] %s\n", hexString)

	reader := bufio.NewReader(bytes.NewReader(msg.Payload()))

	header := &decoder.Header{Buf: reader}
	// if err := header.Validate(); err != nil {
	// 	log.Fatalf("validation failed, %s", err.Error())
	// }

	headerDecoded, err := header.Decode()
	if err != nil {
		log.Fatalf("decoding failed, %s", err.Error())
	}

	fmt.Printf("%+v\n", headerDecoded)

	report := &decoder.Report{Header: headerDecoded, Buf: header.Buf}
	// if err := report.Validate(); err != nil {
	// 	log.Fatalf("validation failed, %s", err.Error())
	// }

	reportDecoded, err := report.Decode()
	if err != nil {
		log.Fatalf("decoding failed, %s", err.Error())
	}

	fmt.Printf("%+v\n", reportDecoded)
}
