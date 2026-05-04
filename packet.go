package main

type Packet struct {
	Type    string
	Payload string
}

const (
	TypeHeartbeat = "HEARTBEAT"
	TypeMessage   = "MESSAGE"
)
