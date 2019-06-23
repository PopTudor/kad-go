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

	c3 := NewNodeIdWith(no1.NodeId.Key)

	no1.RoutingTable.Add(c1)
	no1.RoutingTable.Add(c2)
	no1.RoutingTable.Add(c3)
	//no1.RoutingTable.Describe()
	fmt.Println()

}
