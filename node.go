package main

import (
	"encoding/json"
	"net"
	"time"
)

type Node struct {
	name         string
	address      string
	registryAddr string
	messages     []*Message
}

func (n *Node) SendHeartBeat() {
	for {
		time.Sleep(2 * time.Second)
		conn, err := net.Dial("tcp", n.registryAddr)
		handleError(err)

		// write a json Packet
		packet := Packet{
			Type:    TypeHeartbeat,
			Payload: n.name,
		}

		jsonData, _ := json.Marshal(packet)

		_, err = conn.Write(append(jsonData, '\n'))
		handleError(err)

		conn.Close()
	}
}
