package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

type Consumer struct {
	ring *Ring
}

func NewConsumer(ring *Ring) *Consumer {
	return &Consumer{
		ring: ring,
	}
}

func (c *Consumer) Read(key string, offset int) (Message, error) {
	nodeipaddress := c.ring.GetNode(key)

	conn, err := net.Dial("tcp", nodeipaddress)
	handleError(err)

	fetchRequest := FetchRequest{Offset: offset}

	fetchRequestBytes, _ := json.Marshal(fetchRequest)

	packet := Packet{
		Type:    TypeFetch,
		Key:     key,
		Payload: fetchRequestBytes,
	}

	packetBytes, _ := json.Marshal(packet) // because we need to sent slice of bytes over the wire

	_, err = conn.Write(append(packetBytes, '\n'))
	handleError(err)

	// read the response back from the node
	reader := bufio.NewReader(conn)              // wraps raw tcp connection with ability to say read untilyou hit this char
	responseBytes, err := reader.ReadBytes('\n') // keep reading bytes until you hit \n

	// what does node give back, it has to be a packet right?
	var responsePacket Packet
	err = json.Unmarshal(responseBytes, &responsePacket)
	if err != nil {
		fmt.Println("error", err)
		return Message{}, err
	}

	msgType := responsePacket.Type
	msgPayload := responsePacket.Payload

	var respMessage Message
	if msgType == TypeMessage {
		err = json.Unmarshal(msgPayload, &respMessage)
		handleError(err)
	}

	return respMessage, err
}
