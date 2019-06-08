package main

import (
	"encoding/json"
	"fmt"
)

type MessageType int

const (
	PING       MessageType = 0
	PONG       MessageType = 1
	FIND_NODE  MessageType = 2
	FIND_VALUE MessageType = 3
	STORE      MessageType = 4
)

func (e MessageType) String() string {
	switch e {
	case PING:
		return "PING"
	case PONG:
		return "PONG"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type Message struct {
	Type     MessageType `json:"type"`
	TO       Id          `json:"to"`
	From     Id          `json:"from"`
	FileHash string      `json:"file_hash"`
	Contacts []NodeId    `json:"contacts"`
}

func (m *Message) String() string {
	res2B, _ := json.Marshal(m)
	return fmt.Sprintf("%s", string(res2B))
}
