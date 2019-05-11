package main

import (
	"fmt"
	"net"
)

const DistanceBuckets = 20
const NodesInBucket = 20

type RoutingTable struct {
	currentNode NodeID
	buckets     [DistanceBuckets]Bucket
}

func NewRoutingTable(id NodeID) *RoutingTable {
	return &RoutingTable{
		currentNode: id,
		buckets:     [DistanceBuckets]Bucket{},
	}
}

func (rt *RoutingTable) Add(contact Contact) {
	prefixLen := rt.currentNode.SharedPrefixLen(&contact.ID)
	if prefixLen >= DistanceBuckets {
		// if outside of bucket do what? ignore or return something?
		return
	}
	rt.buckets[prefixLen].Add(&contact)
}
func (rt *RoutingTable) Describe() {
	rt.currentNode.Describe()
	for bucket := range rt.buckets {
		fmt.Printf("Bucket %d [", bucket)
		rt.buckets[bucket].Describe()
		fmt.Println("]")
	}
}

type Bucket struct {
	Contacts []*Contact
}

func (b *Bucket) Describe() {
	for contact := range b.Contacts {
		fmt.Printf("Contact %d: ", contact)
		b.Contacts[contact].Describe()
	}
}

func (b *Bucket) Add(contact *Contact) {
	if len(b.Contacts) >= NodesInBucket {
		// if this happens we should actually ping each node and remove the slowest from the list instead of the last one
		b.Contacts = b.Contacts[1:] // pop back item
	}
	b.Contacts = append([]*Contact{contact}, b.Contacts...) // push front
}

type Contact struct {
	ID NodeID
	IP net.IPAddr
}

func (c *Contact) Describe() {
	fmt.Printf("{%s / %s}", c.ID.String(), c.IP.String())
}
