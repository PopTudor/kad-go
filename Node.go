package main

import (
	"encoding/json"
	"errors"
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
	id := NewNodeKey()
	contact := NewNodeIdWith(id)
	return &Node{
		NodeId:       &contact,
		RoutingTable: NewRoutingTable(&contact),
	}
}
func NewNodeWithId(id NodeId) *Node {
	return &Node{
		NodeId:       &id,
		RoutingTable: NewRoutingTable(&id),
	}
}

func NewNodeWithPort(port uint16) *Node {
	if port > 65535 {
		panic("Port too big")
	}
	id := NewNodeKey()

	address := fmt.Sprintf("127.0.0.1:%d", port)
	ip, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}
	contact := NewNodeIdWithIp(id, ip)
	return &Node{
		NodeId:       &contact,
		RoutingTable: NewRoutingTable(&contact),
	}
}

func NewNodeWithKey(key Key) *Node {
	nodeId := NewNodeIdWith(key)
	n := NewNode()
	n.NodeId = &nodeId
	return n
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
		n.handleConnection(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func (n *Node) handleConnection(conn net.Conn) {
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

	if msg.TO != n.NodeId.key {
		fmt.Println("Ignored. Not targeted node")
		return
	}
	fmt.Printf("%s <<< %s \n", n, msg)

	from := msg.From
	to := msg.TO
	msg.TO = from
	msg.From = to
	switch msg.Type {
	case PING:
		msg.Type = PONG
	case FIND_NODE:
		msg.Bucket = n.RoutingTable.FindClosestBucketById(msg.FindId)
	}

	fmt.Printf("%s >>> %s \n", n, msg)
	encoder.Encode(&msg)
}

// ping a node to find out if is online
func (n *Node) Ping(other *Node) {
	fmt.Println(other.NodeId.IP.String())
	conn, err := net.DialTCP("tcp", nil, other.NodeId.IP)
	checkError(err)

	msg := Message{
		Type: PING,
		From: n.NodeId.key,
		TO:   other.NodeId.key,
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
func (n *Node) FindNode(node NodeId) (*NodeId, error) {
	bucket := n.RoutingTable.FindClosestBucket(&node)
	hasNode, nodeIndex := bucket.Has(node)
	if hasNode {
		get := bucket.Get(nodeIndex)
		return &get, nil
	} else {
		found, err := n.findNodeRemote(node, bucket)
		if err == nil {
			n.RoutingTable.Add(found)
		}
		return found, err
	}
	return nil, errors.New("Node not found")
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

func (n *Node) findNodeRemote(searchedNode NodeId, bucket Bucket) (*NodeId, error) {
	has, _ := bucket.Has(*n.NodeId)
	if has {
		return nil, errors.New("Node not found at remote nodes")
	}

	for _, item := range bucket.nodes {
		conn, err := net.DialTCP("tcp", nil, item.IP)
		checkError(err)

		msg := Message{
			Type:   FIND_NODE,
			From:   n.NodeId.key,
			TO:     item.key,
			FindId: searchedNode.key,
		}
		fmt.Printf("%s >>> %s\n", n, msg)
		encoder := json.NewEncoder(conn)
		decoder := json.NewDecoder(conn)

		encoder.Encode(msg)
		decoder.Decode(&msg)

		fmt.Printf("%s <<< %s\n", n, msg)

		hasNode, index := msg.Bucket.Has(searchedNode)
		if hasNode {
			node := msg.Bucket.Get(index)
			return &node, nil
		} else if !msg.Bucket.IsEmpty() {
			return n.findNodeRemote(searchedNode, msg.Bucket)
		} else {
			continue
		}
	}
	return nil, errors.New("Node not found or not in the network")
}

/**
 *
 */
type FindValueResponse struct {
	ValueFound Segment
	Contacts   []NodeId
}
