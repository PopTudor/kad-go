package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
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

// ping a node to find out if is online
func (n *Node) Ping(other *Node) {
	fmt.Println(other.Contact.IP.String())
	conn, err := net.DialTCP("tcp", nil, other.Contact.IP)
	if err != nil {
		panic(err)
	}

	msg := Message{
		Type: PING,
		From: *n.Contact.ID,
		TO:   *other.Contact.ID,
	}
	fmt.Printf("Ping req: (from %s) (to %s) \n", n.Contact.ID.String(), other.Contact.IP.String())
	pingJson, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	var num = uint16(len(pingJson))

	err = binary.Write(conn, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}
	binary.Write(conn, binary.BigEndian, &pingJson)

	success := false

	err = binary.Read(conn, binary.BigEndian, &success)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ping resp: %t\n", success)

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

	var msgSize uint16
	err := binary.Read(conn, binary.BigEndian, &msgSize)
	if err != nil {
		err = binary.Write(conn, binary.BigEndian, false)
		if err != nil {
			panic(err)
		}
	}

	data := make([]byte, msgSize)
	_, err = io.ReadFull(conn, data)
	if err != nil {
		err = binary.Write(conn, binary.BigEndian, false)
		if err != nil {
			panic(err)
		}
	}

	msg := &Message{}
	err = json.Unmarshal(data, msg)
	if err != nil {
		err = binary.Write(conn, binary.BigEndian, false)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("Recv (from %s / msg %s)\n", msg.From.String(), msg)

	err = binary.Write(conn, binary.BigEndian, true)
	if err != nil {
		panic(err)
	}
}

/**
 *
 */
type FindValueResponse struct {
	ValueFound Segment
	Contacts   []Contact
}
