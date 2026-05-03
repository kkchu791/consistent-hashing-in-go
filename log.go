package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Log struct {
	name   string
	file   *os.File
	offset int
}

func NewLog(name string) *Log {
	// open /create the file named name + ".log"
	// return a log struct with the file handle and offsetarting at 0

	filename := name + ".log"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	handleError(err)

	return &Log{
		name:   name,
		file:   f,
		offset: 0,
	}
}

func (l *Log) Append(m *Message) error {
	// serialize the message to JSON
	jsonData, err := json.Marshal(m) //stuct to -> {"task": "sup"}
	handleError(err)
	// Writes it as a new line at the end of the file
	_, err = l.file.Write(append(jsonData, '\n'))
	handleError(err)
	// incremement the offset counter
	l.offset++

	return nil
}

func (l *Log) Read(offset int) (Message, error) {
	l.file.Seek(0, 0)
	scanner := bufio.NewScanner(l.file)
	var m Message
	for i := 0; scanner.Scan(); i++ {
		if i == offset {
			json.Unmarshal([]byte(scanner.Text()), &m)
			return m, nil
		}
	}

	return m, fmt.Errorf("offset %d not found", offset)

}
