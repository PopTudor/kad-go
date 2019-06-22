package main

import (
	"fmt"
	"math/rand"
	"net"
)

type NodeId struct {
	key Key
	IP  *net.TCPAddr
}

func newRandomPort() int {
	port := rand.Intn(65535)
	return port
}

func NewNodeId() NodeId {
	port := newRandomPort()
	address := fmt.Sprintf("127.0.0.1:%d", port)
	ip, err := net.ResolveTCPAddr("tcp", address)

	if err != nil {
		panic(err)
	}
	id := NewNodeKey()
	return NewNodeIdWithIp(id, ip)
}

func NewNodeIdWith(id Key) NodeId {
	n := NewNodeId()
	n.key = id
	return n
}

func NewNodeIdWithIp(id Key, addr *net.TCPAddr) NodeId {
	return NodeId{key: id, IP: addr}
}

func (c *NodeId) DistanceTo(id *NodeId) uint32 {
	return c.key.SharedPrefixLen(id.key)
}

func (c *NodeId) Describe() {
	fmt.Printf("%s", c)
}
func (c *NodeId) String() string {
	return fmt.Sprintf("{%s / %s}", c.key.StringHex(), c.IP.String())
}
