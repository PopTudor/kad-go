package main

import (
	"encoding/json"
	"fmt"
)

type MessageType int
type ErrorType int

const (
	NOT_FOUND ErrorType = 0
)

const (
	PING       MessageType = 0
	PONG       MessageType = 1
	FIND_NODE  MessageType = 2
	FIND_VALUE MessageType = 3
	STORE      MessageType = 4
	ERROR      MessageType = -1
)

func (e MessageType) String() string {
	switch e {
	case PING:
		return "PING"
	case PONG:
		return "PONG"
	case FIND_NODE:
		return "FIND_NODE"
	case FIND_VALUE:
		return "FIND_VALUE"
	case STORE:
		return "STORE"
	case ERROR:
		return "ERROR"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type Message struct {
	Type     MessageType `json:"type"`
	From     Key         `json:"from"`
	TO       Key         `json:"to"`
	FileHash string      `json:"file_hash"`
	Nodes    []NodeId    `json:"nodes"`
	FindId   Key         `json:"find_id"`
}

func (m *Message) String() string {
	res2B, _ := json.Marshal(m)
	return fmt.Sprintf("%s", string(res2B))
}

func (m *Message) Has(id NodeId) (bool, int16) {
	if m.Nodes == nil {
		return false, -1
	}
	for i, contact := range m.Nodes {
		if contact.Key == id.Key {
			return true, int16(i)
		}
	}
	return false, -1
}