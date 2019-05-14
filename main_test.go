package main

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

func TestBasic(t *testing.T) {
	n1 := NewNodeID()
	n2 := NewNodeID()
	n1.DescribeBinary()
	n2.DescribeBinary()

	len1 := n1.SharedPrefixLen(&n2)
	len2 := n2.SharedPrefixLen(&n1)
	if len1 != len2 {
		t.Error("Commutative property has failed. A+B = B+A")
	}
	fmt.Printf("%d", len1)
	ip, _ := net.ResolveTCPAddr("tcp", ":5433")
	no1 := NewNode()
	c1 := NewContactWith(&n1)
	ids := NewNodeID()
	c2 := NewContactWithIp(&ids, ip)

	c3 := NewContactWith(no1.Contact.ID)

	no1.RoutingTable.Add(*c1)
	no1.RoutingTable.Add(*c2)
	no1.RoutingTable.Add(*c3)
	//no1.RoutingTable.Describe()
	fmt.Println()

}

func TestNode_Ping(t *testing.T) {
	n1 := NewNode()
	n2 := NewNode()
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		n1.Start()
	}()
	go func() {
		n2.Ping(n1)
		w.Done()
	}()
	w.Wait()
}
