package main

import (
	"fmt"
	"testing"
)

func TestRoutingTable_AddToEnds(t *testing.T) {
	c := NewNodeId()
	rt := NewRoutingTable(&c)

	from := NewKeyFrom("6C7D63826DE1F6529E4E248771CA45FB69CC397B")
	nc := NewNodeIdWith(from)
	nc.Describe()
	index := rt.Add(&nc)
	fmt.Printf("addet at index: %d\n", index)
	index = rt.Add(&c)
	fmt.Printf("addet at index: %d\n", index)
}
func TestNode_At_Correct_Bucket(t *testing.T) {
	ka := NewKeyFrom("")
	a := NewNodeWithKey(ka)
	t.Log(a.String())

	// a should have itself in bucket 0
	if !a.RoutingTable.IsNodeIdInBucket(a.NodeId, 0) {
		t.Fatal("Node not in desired bucket")
	}
	// node at index 1 in a's routing table
	kb := NewKeyFrom("00000000000000000001")
	t.Log(kb.String())
	b := NewNodeWithKey(kb)

	// a should be in bucket 0
	// b should be in bucket 1
	a.RoutingTable.Add(b.NodeId)

	b.RoutingTable.Add(a.NodeId)

}
