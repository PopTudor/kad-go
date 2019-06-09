package main

import (
	"fmt"
	"net"
)

type NodeId struct {
	ID Id
	IP *net.TCPAddr
}

func NewContact() *NodeId {
	ip, err := net.ResolveTCPAddr("tcp", ":5443")
	if err != nil {
		panic(err)
	}
	id := NewNodeID()
	return NewContactWithIp(id, ip)
}

func NewContactWith(id Id) *NodeId {
	return &NodeId{ID: id}
}

func NewContactWithIp(id Id, addr *net.TCPAddr) *NodeId {
	return &NodeId{ID: id, IP: addr}
}

func (c *NodeId) DistanceTo(id NodeId) uint32 {
	return c.ID.SharedPrefixLen(id.ID)
}

func (c *NodeId) Describe() {
	fmt.Printf("%s", c)
}
func (c *NodeId) String() string {
	return fmt.Sprintf("{%s / %s}", c.ID.StringHex(), c.IP.String())
}
