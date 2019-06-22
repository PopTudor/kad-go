package main

import (
	"fmt"
	"net"
	"testing"
)

func TestBasic(t *testing.T) {
	n1 := NewNodeKey()
	n2 := NewNodeKey()
	n1.DescribeBinary()
	n2.DescribeBinary()

	len1 := n1.SharedPrefixLen(n2)
	len2 := n2.SharedPrefixLen(n1)
	if len1 != len2 {
		t.Error("Commutative property has failed. A+B = B+A")
	}
	fmt.Printf("%d", len1)
	ip, _ := net.ResolveTCPAddr("tcp", ":5433")
	no1 := NewNode()
	c1 := NewNodeIdWith(n1)
	ids := NewNodeKey()
	c2 := NewNodeIdWithIp(ids, ip)

	c3 := NewNodeIdWith(no1.NodeId.key)

	no1.RoutingTable.Add(&c1)
	no1.RoutingTable.Add(&c2)
	no1.RoutingTable.Add(&c3)
	//no1.RoutingTable.Describe()
	fmt.Println()

}

func TestNode_Valid_NodeId_Ping(t *testing.T) {
	n1 := NewNode()
	n2 := NewNode()
	go n1.Start()

	n2.Ping(n1)
}
func TestNode_InValid_NodeId_Ping(t *testing.T) {
	n1 := NewNode()
	n2 := NewNode()
	n3 := NewNode()
	go n1.Start()
	go n2.Start()
	go n3.Start()

	n1.Ping(n3)
}
func TestNode_FindNode(t *testing.T) {
	n1 := NewNode()
	n2 := NewNode()
	go n1.Start()
	go n2.Start()
	n1.RoutingTable.Add(n2.NodeId)
	_, ok := n1.FindNode(*n2.NodeId)
	if ok != nil {
		panic("Nodeid not found in routing table")
	}
}
func TestNode_FindNode_Network(t *testing.T) {
	n1 := NewNode()
	go n1.Start()

	n2 := NewNode()
	go n2.Start()

	n3 := NewNode()
	go n3.Start()

	n1.RoutingTable.Add(n2.NodeId)
	n3.RoutingTable.Add(n1.NodeId)

	foundNode, err := n3.FindNode(*n2.NodeId)
	if err != nil {
		panic("Nodeid not found in routing table")
	}
	if foundNode.key != n2.NodeId.key {
		panic("found wrong node")
	}
}

func TestNode_FindNodeRecursive(t *testing.T) {
	// current node at index 0 in the routing table
	a := NewNodeWithKey(NewKeyFrom(""))
	t.Log(a.String())
	go a.Start()

	// node at index 1 in a's routing table
	k1 := NewKeyFrom("00000000000000000001")
	t.Log(k1.String())
	n1 := NewNodeWithKey(k1)
	go n1.Start()

	a.RoutingTable.Add(n1.NodeId)

	for i := 0; i > 110; i-- {
		n := NewNodeWithKey(NewKeyFrom(fmt.Sprintf("%d", i)))
		go n.Start()
		n1.RoutingTable.Add(n.NodeId)
	}

	lastBucket := n1.RoutingTable.LastBucket()
	lastNode := lastBucket.LastNode()

	if found, err := a.FindNode(lastNode); err == nil {
		if found.key != lastNode.key {
			t.Fatal("Found wrong node")
		}
	} else {
		t.Fatal("Last node not found ")
	}

	firstNode := lastBucket.Get(0)
	if found, err := a.FindNode(firstNode); err == nil {
		if found.key != firstNode.key {
			t.Fatal("found wrong node")
		}
	} else {
		t.Fatal(err)
	}
}

func NewNodeAtDistance(node *Node, distance uint32) Node {
	nodeid := NewNodeId()
	for {
		distanceTo := nodeid.DistanceTo(node.NodeId)
		if distanceTo == distance {
			node1 := NewNodeWithId(nodeid)
			return *node1
		}
		nodeid = NewNodeId()
	}
}
