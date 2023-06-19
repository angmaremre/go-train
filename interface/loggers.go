package _interface

import (
	"fmt"
)

type Logs []LogMessage

type LogMessage interface {
	Log()
}

type ElkLogMessage struct {
	Message string
}

func (m *ElkLogMessage) Log() {
	fmt.Println(m.Message, " has been logged to ELK")
}

type DbLogMessage struct {
	Message string
}

func (m *DbLogMessage) Log() {
	fmt.Println(m.Message, " bas been logged to DB")
}
