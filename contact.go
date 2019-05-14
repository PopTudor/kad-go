package main

import (
	"fmt"
	"net"
)

type Contact struct {
	ID *NodeID
	IP *net.TCPAddr
}

func NewContact() *Contact {
	ip, err := net.ResolveTCPAddr("tcp", ":5443")
	if err != nil {
		panic(err)
	}
	id := NewNodeID()
	return NewContactWithIp(&id, ip)
}
func NewContactWith(id *NodeID) *Contact {
	return &Contact{ID: id}
}

func NewContactWithIp(id *NodeID, addr *net.TCPAddr) *Contact {
	return &Contact{ID: id, IP: addr}
}

func (c *Contact) Describe() {
	fmt.Printf("{%s / %s}", c.ID.String(), c.IP.String())
}
