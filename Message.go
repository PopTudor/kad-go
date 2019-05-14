package main

import (
	"encoding/json"
	"fmt"
)

type MessageType int

const (
	PING       MessageType = 0
	FIND_NODE  MessageType = 1
	FIND_VALUE MessageType = 2
	STORE      MessageType = 3
)

type Message struct {
	Type     MessageType `json:"type"`
	From     NodeID      `json:"from"`
	TO       NodeID      `json:"to"`
	FileHash string      `json:"file_hash"`
	Contacts []Contact   `json:"contacts"`
}

func (m *Message) String() string {
	res2B, _ := json.Marshal(m)
	return fmt.Sprintf("%s", string(res2B))
}
