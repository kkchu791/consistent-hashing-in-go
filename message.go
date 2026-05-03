package main

import "time"

type Message struct {
	Task      string
	Notes     string
	Timestamp time.Time
}

func NewMessage(m, n string) *Message {
	return &Message{
		Task:      m,
		Notes:     n,
		Timestamp: time.Now(),
	}
}
