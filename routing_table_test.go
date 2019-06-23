package main

import (
	"testing"
)

func Test_Add_Node_At_Correct_Bucket(t *testing.T) {
	ka := NewKeyFrom([]byte{})
	a := NewNodeWithKey(ka)
	t.Log(a.NodeId.key.String())

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

	kn := NewKeyAtIndexWithBitsSetTo(255, 0)
	n := NewNodeWithKey(kn)
	t.Log(n.NodeId.key.String())

	a.RoutingTable.Add(n.NodeId)

	if !a.RoutingTable.IsNodeIdInBucket(n.NodeId, Distance_Buckets-1) {
		t.Fatal("Node n should be in last bucket")
	}

}
