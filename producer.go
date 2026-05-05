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

	// message looks like {"task": "Jog", "notes": "2 miles"}
	messageBytes, _ := json.Marshal(message)

	packet := Packet{
		Type:    TypeMessage,
		Key:     key,
		Payload: messageBytes,
	}

	packetBytes, _ := json.Marshal(packet)
	_, err = conn.Write(packetBytes)
	handleError(err)

	return err
}
