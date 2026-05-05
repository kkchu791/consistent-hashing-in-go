package main

import "encoding/json"

type Packet struct {
	Type    string
	Key     string
	Payload json.RawMessage
}

const (
	TypeHeartbeat = "HEARTBEAT"
	TypeMessage   = "MESSAGE"
)
