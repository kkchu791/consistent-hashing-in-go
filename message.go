package main

import "time"

type Message struct {
	task        string
	notes       string
	improvement string
	goal        string
	timestamp   time.Time
	offset      int
}
