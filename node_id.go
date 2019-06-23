package main

import (
	"fmt"
	"math/rand"
	"net"
)

type NodeId struct {
	Key Key
	IP  *net.TCPAddr
}

func newRandomPort() int {
	port := rand.Intn(65535)
	return port
}

func NewNodeId() NodeId {
	key := NewNodeKey()
	return NewNodeIdWith(key)
}

func NewNodeIdWith(key Key) NodeId {
	port := newRandomPort()
	address := fmt.Sprintf("127.0.0.1:%d", port)
	ip, err := net.ResolveTCPAddr("tcp", address)

	if err != nil {
		panic(err)
	}
	return NewNodeIdWithIp(key, ip)
}

func NewNodeIdWithIp(id Key, addr *net.TCPAddr) NodeId {
	return NodeId{Key: id, IP: addr}
}

func (c *NodeId) DistanceTo(id *NodeId) uint32 {
	return c.Key.SharedPrefixLen(id.Key)
}

func (c *NodeId) Describe() {
	fmt.Printf("%s", c)
}
func (c *NodeId) String() string {
	return fmt.Sprintf("{%s / %s}", c.Key.StringHex(), c.IP.String())
}
