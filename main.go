package main

import (
	"fmt"
	"os"
)

func main() {
	ring := NewRing(3)
	registry := NewRegistry("127.0.0.1:3000")

	node1 := Node{name: "node1", address: "127.0.0.1:8080", registryAddr: "127.0.0.1:3000"}
	node2 := Node{name: "node2", address: "127.0.0.1:8081", registryAddr: "127.0.0.1:3000"}
	node3 := Node{name: "node3", address: "127.0.0.1:8082", registryAddr: "127.0.0.1:3000"}

	node1.Listen()
	node2.Listen()
	node3.Listen()

	registry.Register(node1)
	registry.Register(node2)
	registry.Register(node3)

	ring.AddNode(node1)
	ring.AddNode(node2)
	ring.AddNode(node3)

	newMessage := NewMessage(
		"Learning",
		"Learning Distribute Systems",
	)

	np := NewProducer(ring)

	np.Send("distributed-systems", newMessage)

}

func handleError(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

// go registry.ListenHeartBeat()

// go node1.SendHeartBeat()
// go node2.SendHeartBeat()
// go node3.SendHeartBeat()

// fmt.Println(ring.GetNode("lakers_score"))
// fmt.Println(ring.GetNode("rockets_score"))
// fmt.Println(ring.GetNode("celtics_score"))
// fmt.Println("before check health:", ring.GetNode("lakers_score"))

// go registry.StartHealthChecker(ring)
// // r.RemoveNode(node2)

// time.Sleep(20 * time.Second)

// fmt.Println("after check health:", ring.GetNode("lakers_score"))

// // testing log storage
// l := NewLog("distributed-systems")
// m := NewMessage(
// 	"Implement log storage",
// 	"Learning how to test log storage append and read methods",
// )
// l.Append(m)
// readmessage, err := l.Read(0) // what if I wanted to Read certain message with "task: Log Implementation"
// handleError(err)
// fmt.Println(readmessage)
