package main

import (
	"fmt"
)

const DistanceBuckets = 160

type RoutingTable struct {
	currentNode *Contact
	// buckets with index closer to 0 store contacts further from the current node because they share less prefix bits
	// the current node is in the last bucket because shared prefix len is 160
	buckets [DistanceBuckets]Bucket
}

func NewRoutingTable(id *Contact) *RoutingTable {
	return &RoutingTable{
		currentNode: id,
		buckets:     [DistanceBuckets]Bucket{},
	}
}

func (rt *RoutingTable) Add(contact Contact) {
	prefixLen := rt.currentNode.ID.SharedPrefixLen(contact.ID)
	if prefixLen == DistanceBuckets {
		rt.buckets[prefixLen-1].Add(&contact)
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
