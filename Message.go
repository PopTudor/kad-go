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
	From     Id          `json:"from"`
	TO       Id          `json:"to"`
	FileHash string      `json:"file_hash"`
	Bucket   Bucket      `json:"bucket"`
	FindId   Id          `json:"find_id"`
}

func (m *Message) String() string {
	res2B, _ := json.Marshal(m)
	return fmt.Sprintf("%s", string(res2B))
}
