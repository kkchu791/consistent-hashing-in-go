package main

import (
	"bufio"
	"encoding/json"
	"fmt"
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

		nameBytes, _ := json.Marshal(n.name)

		packet := Packet{
			Type:    TypeHeartbeat,
			Key:     "registry",
			Payload: json.RawMessage(nameBytes),
		}

		jsonData, _ := json.Marshal(packet)

		_, err = conn.Write(append(jsonData, '\n'))
		handleError(err)

		conn.Close()
	}
}

func (n *Node) Listen() {
	listener, err := net.Listen("tcp", n.address)
	handleError(err)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		handleError(err)

		go handleNodeConnection(conn)
	}
}

// reads from the incoming and writes to the log file
func handleNodeConnection(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		packetBytes, err := reader.ReadBytes('\n')

		if err != nil {
			fmt.Println("error", err)
			return
		}

		var packet Packet

		err = json.Unmarshal(packetBytes, &packet) // there's an assumption that unmarshalling will read the data to the correct fields of the packet to create the json
		if err != nil {
			fmt.Println("error", err)
			return
		}

		// write the packet to the log
		// before that we have to extract type and message from the packet
		msgType := packet.Type // "TypeMessage"
		msgKey := packet.Key
		msgPayloadBytes := packet.Payload

		if msgType == TypeMessage {
			message := NewMessage("", "")
			err = json.Unmarshal(msgPayloadBytes, message)
			if err != nil {
				fmt.Println("error", err)
				return
			}

			l := NewLog(msgKey)
			l.Append(message)
		} else {
			fmt.Println("got some other type other than messagetype")
		}
	}
}
