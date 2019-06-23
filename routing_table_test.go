package main

import (
	"testing"
)

func setupNodeZero(t *testing.T) *Node {
	ka := NewKeyFrom([]byte{})
	a := NewNodeWithKey(ka)
	t.Log(a.NodeId.key.String())
	return a
}

func Test_Add_Node_At_Beginning_Of_Bucket(t *testing.T) {
	a := setupNodeZero(t)

	// a should have itself in bucket 0
	if !a.RoutingTable.IsNodeIdInBucket(a.NodeId, 0) {
		t.Fatal("Node not in desired bucket")
	}
	// node at index 1 in a's routing table
	kb := NewKeyAtIndexWithBitsSetTo(1, KeyLen-1)
	b := NewNodeWithKey(kb)
	t.Log(b.NodeId.key.String())

	// a should be in bucket 0
	// b should be in bucket 1
	a.RoutingTable.Add(b.NodeId)
	if !a.RoutingTable.IsNodeIdInBucket(b.NodeId, 1) {
		t.Fatal("Node b should have been in bucket 2")
	}
}

func Test_Add_NodeId_At_End_Of_Bucket(t *testing.T) {
	a := setupNodeZero(t)

	kn := NewKeyAtIndexWithBitsSetTo(255, 0)
	n := NewNodeWithKey(kn)
	t.Log(n.NodeId.key.String())

	a.RoutingTable.Add(n.NodeId)

	if !a.RoutingTable.IsNodeIdInBucket(n.NodeId, 159) {
		t.Fatal("Node n should be in last bucket")
	}

	// node id 11111110 should be added to bucket 159 because starts with 1
	kn159 := NewKeyAtIndexWithBitsSetTo(254, 0)
	n159 := NewNodeWithKey(kn159)
	t.Log(n159.NodeId.key.String())
	a.RoutingTable.Add(n159.NodeId)

	if !a.RoutingTable.IsNodeIdInBucket(n159.NodeId, 159) {
		t.Fatal("Node n should be in last bucket")
	}

	// node id 01111110 should be added to bucket 158 because starts with 1
	kn158 := NewKeyAtIndexWithBitsSetTo(63, 0)
	n158 := NewNodeWithKey(kn158)
	t.Log(n158.NodeId.key.String())
	a.RoutingTable.Add(n158.NodeId)

	if !a.RoutingTable.IsNodeIdInBucket(n158.NodeId, 158) {
		t.Fatal("Node n158 should be in n-1 bucket eg 158")
	}
	// byte 00111111 should add the node at Distance_Buckets - 2
	kn157 := NewKeyAtIndexWithBitsSetTo(31, 0)
	n157 := NewNodeWithKey(kn157)
	t.Log(n157.NodeId.key.String())
	a.RoutingTable.Add(n157.NodeId)

	if !a.RoutingTable.IsNodeIdInBucket(n157.NodeId, 157) {
		t.Fatal("Node n158 should be in n-1 bucket eg 158")
	}
}
