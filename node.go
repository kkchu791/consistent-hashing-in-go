package main

import (
	"net"
	"time"
)

type Node struct {
	name         string
	address      string
	registryAddr string
}

func (n *Node) SendHeartBeat() {
	for {
		time.Sleep(2 * time.Second)
		conn, err := net.Dial("tcp", n.registryAddr)
		handleError(err)

		_, err = conn.Write([]byte(n.name + "\n"))
		handleError(err)

		conn.Close()
	}
}
