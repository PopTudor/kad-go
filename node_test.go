package main

import "testing"

func TestNode_Valid_NodeId_Ping(t *testing.T) {
	n1 := NewNode()
	n2 := NewNode()
	n1.Start()

	n2.Ping(n1)
}
func TestNode_InValid_NodeId_Ping(t *testing.T) {
	n1 := NewNode()
	n2 := NewNode()
	n3 := NewNode()
	n1.Start()
	n2.Start()
	n3.Start()

	msg := n1.Ping(n3)
	if msg.Type != PONG {
		t.Error("Message is not PONG")
	}
	if msg.From != n3.NodeId.Key {
		t.Error("PONG is not from desired node")
	}
}

func TestNode_FindNode(t *testing.T) {
	n1 := NewNode()
	n2 := NewNode()
	go n1.Start()
	go n2.Start()
	n1.RoutingTable.Add(n2.NodeId)
	_, ok := n1.FindNode(n2.NodeId)
	if ok != nil {
		panic("Nodeid not found in routing table")
	}
}
func TestNode_FindNode_Network(t *testing.T) {
	n1 := NewNode()
	n1.Start()

	n2 := NewNode()
	n2.Start()

	n3 := NewNode()
	n3.Start()

	n1.RoutingTable.Add(n2.NodeId)
	n3.RoutingTable.Add(n1.NodeId)

	foundNode, err := n3.FindNode(n2.NodeId)
	if err != nil {
		// n3 entered recursive mode with itself
		panic(err)
	}
	if foundNode.Key != n2.NodeId.Key {
		panic("found wrong node")
	}
}

func TestNode_FindNodeRecursive(t *testing.T) {
	// current node at index 0 in the routing table
	//a := NewNodeWithKey(NewKeyFrom(""))
	//t.Log(a.String())
	//go a.Start()
	//
	//node at index 1 in a's routing table
	//k1 := NewKeyFrom("00000000000000000001")
	//t.Log(k1.String())
	//n1 := NewNodeWithKey(k1)
	//go n1.Start()

	//a.RoutingTable.Add(n1.NodeId)

	//for i := 0; i > 110; i-- {
	//	n := NewNodeWithKey(NewKeyFrom(fmt.Sprintf("%d", i)))
	//	go n.Start()
	//	n1.RoutingTable.Add(n.NodeId)
	//}

	//lastBucket := n1.RoutingTable.LastBucket()
	//lastNode := lastBucket.LastNode()
	//
	//if found, err := a.FindNode(lastNode); err == nil {
	//	if found.key != lastNode.key {
	//		t.Fatal("Found wrong node")
	//	}
	//} else {
	//	t.Fatal("Last node not found ")
	//}
	//
	//firstNode := lastBucket.Get(0)
	//if found, err := a.FindNode(firstNode); err == nil {
	//	if found.key != firstNode.key {
	//		t.Fatal("found wrong node")
	//	}
	//} else {
	//	t.Fatal(err)
	//}
}
