package main

import (
	"fmt"
	"math/rand"
	"net"
)

type NodeId struct {
	ID Id
	IP *net.TCPAddr
}

func newRandomPort() int {
	port := rand.Intn(65535) + 10.000
	return port
}

func NewContact() *NodeId {
	port := newRandomPort()
	address := fmt.Sprintf("127.0.0.1:%d", port)
	ip, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}
	id := NewNodeID()
	return NewContactWithIp(id, ip)
}

func NewContactWith(id Id) *NodeId {
	n := NewContact()
	n.ID = id
	return n
}

func NewContactWithIp(id Id, addr *net.TCPAddr) *NodeId {
	return &NodeId{ID: id, IP: addr}
}

func (c *NodeId) DistanceTo(id *NodeId) uint32 {
	return c.ID.SharedPrefixLen(id.ID)
}

func (c *NodeId) Describe() {
	fmt.Printf("%s", c)
}
func (c *NodeId) String() string {
	return fmt.Sprintf("{%s / %s}", c.ID.StringHex(), c.IP.String())
}
