package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Node struct {
	NodeId       *NodeId
	RoutingTable *RoutingTable
	DHT          DHT
}

func NewNode() *Node {
	id := NewNodeID()
	ip, err := net.ResolveTCPAddr("tcp", "127.0.0.1:5443")
	if err != nil {
		panic(err)
	}
	contact := NewContactWithIp(&id, ip)
	return &Node{
		NodeId:       contact,
		RoutingTable: NewRoutingTable(contact),
	}
}

func NewNodeWithId(id Id) *Node {
	contact := NewContactWith(&id)
	return &Node{
		NodeId:       contact,
		RoutingTable: NewRoutingTable(contact),
	}
}
func (n *Node) Start() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", n.NodeId.IP.String())
	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
		// handle error
	}
	fmt.Printf("start %s\n", n)
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			panic(err)
		}
		handleConnection(n, conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func handleConnection(n *Node, conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	msg := Message{}
	decoder.Decode(&msg)

	if msg.TO != *n.NodeId.ID {
		fmt.Println("Ignored. Not targeted node")
		return
	}

	fmt.Printf("%s <<< %s \n", n, msg)

	from := msg.From
	to := msg.TO
	msg.TO = from
	msg.From = to
	msg.Type = PONG
	encoder.Encode(&msg)
}

// ping a node to find out if is online
func (n *Node) Ping(other *Node) {
	fmt.Println(other.NodeId.IP.String())
	conn, err := net.DialTCP("tcp", nil, other.NodeId.IP)
	checkError(err)

	msg := Message{
		Type: PING,
		From: *n.NodeId.ID,
		TO:   *other.NodeId.ID,
	}
	fmt.Printf("%s >>> %s\n", n, msg)
	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	encoder.Encode(msg)
	decoder.Decode(&msg)

	fmt.Printf("%s <<< %s\n", n, msg)
}

// call to find a specific node with given id. The recipiend of this call
// looks in it's own routing table and returns a set of contacts that are closeset to
// the NodeId that is being looked up
func (n *Node) FindNode(id Id) []NodeId {
	return nil
}

// this call tries to find a specific file NodeId to be located. If the receiving
// node finds this NodeId in it's own DHT segment, it will return the corresponding
// URL. If not, the recipient node returns a list of contacts that are closest
// to the file NodeId
func (n *Node) FindValue(value []byte) *FindValueResponse {
	return nil
}

// This call is used to store a key/value pair(fileID,location) in the DHT segment of the recipient node
// Upon each successful RPC, both the sending/receiving node insert/update each other's contact info in their
// own routing table
func (n *Node) Store(value FileID, contact NodeId) {

}

func (n *Node) String() string {
	return fmt.Sprintf("%s", n.NodeId)
}

/**
 *
 */
type FindValueResponse struct {
	ValueFound Segment
	Contacts   []NodeId
}
