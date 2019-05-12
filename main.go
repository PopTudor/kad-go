package main

import (
	"fmt"
	"net"
)

func main() {
	n1 := NewNodeID()
	n2 := NewNodeID()
	n1.DescribeBinary()
	n2.DescribeBinary()

	len1 := n1.SharedPrefixLen(&n2)
	fmt.Printf("%d", len1)

	no1 := NewNode()
	c1 := Contact{n1, net.IPAddr{}}
	c2 := Contact{NewNodeID(), net.IPAddr{}}
	c3 := Contact{no1.ID, net.IPAddr{}}

	no1.RoutingTable.Add(c1)
	no1.RoutingTable.Add(c2)
	no1.RoutingTable.Add(c3)
	no1.RoutingTable.Describe()

}
