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
