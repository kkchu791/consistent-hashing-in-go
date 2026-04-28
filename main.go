package main

import "fmt"

func main() {
	ring := NewRing(3)
	registry := NewRegistry()

	node1 := Node{name: "node1", address: "192.168.1.1:8080"}
	node2 := Node{name: "node2", address: "192.168.1.2:8080"}
	node3 := Node{name: "node3", address: "192.168.1.3:8080"}

	registry.Register(node1)
	registry.Register(node2)
	registry.Register(node3)

	ring.AddNode(node1)
	ring.AddNode(node3)

	fmt.Println(ring.GetNode("lakers_score"))
	fmt.Println(ring.GetNode("rockets_score"))
	fmt.Println(ring.GetNode("celtics_score"))
	fmt.Println("before check health:", ring.GetNode("lakers_score"))

	registry.CheckHealth(ring)
	// r.RemoveNode(node2)

	fmt.Println("after check health:", ring.GetNode("lakers_score"))

}
