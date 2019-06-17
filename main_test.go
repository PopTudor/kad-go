package main

import (
	"fmt"
	"net"
	"testing"
)

func TestBasic(t *testing.T) {
	n1 := NewNodeID()
	n2 := NewNodeID()
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
	c1 := NewContactWith(n1)
	ids := NewNodeID()
	c2 := NewContactWithIp(ids, ip)

	c3 := NewContactWith(no1.NodeId.ID)

	no1.RoutingTable.Add(c1)
	no1.RoutingTable.Add(c2)
	no1.RoutingTable.Add(c3)
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
	if foundNode.ID != n2.NodeId.ID {
		panic("found wrong node")
	}
}

func TestNode_FindNodeRecursive(t *testing.T) {
	a := NewNodeWithId(NewNodeIdFrom("0"))
	go a.Start()

	n1 := NewNode()
	go n1.Start()

	a.RoutingTable.Add(n1.NodeId)

	for i := 0; i < 128/8; i++ {
		n := NewNode()
		go n.Start()
		n1.RoutingTable.Add(n.NodeId)
	}

	lastBucket := n1.RoutingTable.LastBucket()
	lastNode := lastBucket.LastNode()

	if found, err := a.FindNode(lastNode); err == nil {
		if found.ID != lastNode.ID {
			t.Fatal("Found wrong node")
		}
	} else {
		t.Fatal("Last node not found ")
	}

	firstNode := lastBucket.Get(0)
	if found, err := a.FindNode(firstNode); err == nil {
		if found.ID != firstNode.ID {
			t.Fatal("found wrong node")
		}
	} else {
		t.Fatal(err)
	}
}
