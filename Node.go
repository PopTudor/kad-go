package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Node struct {
	Contact      *Contact
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
		Contact:      contact,
		RoutingTable: NewRoutingTable(contact),
	}
}

func NewNodeWithId(id NodeID) *Node {
	contact := NewContactWith(&id)
	return &Node{
		Contact:      contact,
		RoutingTable: NewRoutingTable(contact),
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

// ping a node to find out if is online
func (n *Node) Ping(other *Node) {
	fmt.Println(other.Contact.IP.String())
	conn, err := net.DialTCP("tcp", nil, other.Contact.IP)
	checkError(err)

	msg := Message{
		Type: PING,
		From: *n.Contact.ID,
		TO:   *other.Contact.ID,
	}
	fmt.Printf("Ping req: (from %s) (to %s) \n", n.Contact.ID.String(), other.Contact.IP.String())
	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	encoder.Encode(msg)
	decoder.Decode(&msg)

	fmt.Printf("Ping resp: %v\n", msg)
}

// call to find a specific node with given id. The recipiend of this call
// looks in it's own routing table and returns a set of contacts that are closeset to
// the Contact that is being looked up
func (n *Node) FindNode(id NodeID) []Contact {
	return nil
}

// this call tries to find a specific file Contact to be located. If the receiving
// node finds this Contact in it's own DHT segment, it will return the corresponding
// URL. If not, the recipient node returns a list of contacts that are closest
// to the file Contact
func (n *Node) FindValue(value []byte) *FindValueResponse {
	return nil
}

// This call is used to store a key/value pair(fileID,location) in the DHT segment of the recipient node
// Upon each successful RPC, both the sending/receiving node insert/update each other's contact info in their
// own routing table
func (n *Node) Store(value FileID, contact Contact) {

}
func (n *Node) Start() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", n.Contact.IP.String())
	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
		// handle error
	}
	fmt.Println("start")
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			panic(err)
		}
		handleConnection(n, conn)
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

	msg := &Message{}
	decoder.Decode(&msg)

	fmt.Printf("Recv (from %s / msg %s)\n", msg.From, msg.Type)

	from := msg.From
	to := msg.TO
	msg.TO = from
	msg.From = to
	msg.Type = PONG
	encoder.Encode(&msg)
}

/**
 *
 */
type FindValueResponse struct {
	ValueFound Segment
	Contacts   []Contact
}
