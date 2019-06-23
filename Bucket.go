package main

import "fmt"

const NodesInBucket = 20

type Bucket struct {
	nodes []NodeId
}

func NewBucket(ids []NodeId) (Bucket) {
	b := Bucket{
		nodes: ids,
	}
	return b
}

func (b *Bucket) Describe() {
	for node := range b.nodes {
		fmt.Printf("NodeId %d: ", node)
		b.nodes[node].Describe()
	}
}

func (b *Bucket) Add(contact NodeId) {
	if len(b.nodes) >= NodesInBucket {
		// if this happens we should actually ping each node and remove the slowest from the list instead of the last one
		b.Pop()
	}
	b.PushFront(contact)
}

// pop back item
func (b *Bucket) Pop() {
	b.nodes = b.nodes[1:]
}

func (b *Bucket) PushFront(contact NodeId) {
	b.nodes = append([]NodeId{contact}, b.nodes...)
}

func (b *Bucket) Has(id NodeId) (bool, int16) {
	if b.nodes == nil {
		return false, -1
	}
	for i, contact := range b.nodes {
		if contact.Key == id.Key {
			return true, int16(i)
		}
	}
	return false, -1
}

func (b *Bucket) Get(i int16) NodeId {
	return b.nodes[i]
}

func (b *Bucket) IsEmpty() bool {
	return len(b.nodes) == 0
}

func (b *Bucket) LastNode() NodeId {
	return b.nodes[len(b.nodes)-1]
}
