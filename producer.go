package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Producer struct {
	ring *Ring
}

func NewProducer(ring *Ring) *Producer {
	return &Producer{
		ring: ring,
	}
}

func (p *Producer) Send(key string, message *Message) error {
	nodeipaddress := p.ring.GetNode(key)

	fmt.Println(nodeipaddress)
	conn, err := net.Dial("tcp", nodeipaddress)
	handleError(err)

	// message looks like {"task": "Drink Water", "notes": "N/A"}
	messageBytes, _ := json.Marshal(message) // converts json into a slice of bytes

	packet := Packet{
		Type:    TypeMessage,
		Key:     key,
		Payload: messageBytes,
	}

	packetBytes, _ := json.Marshal(packet)
	_, err = conn.Write(append(packetBytes, '\n'))
	handleError(err)

	return err
}
