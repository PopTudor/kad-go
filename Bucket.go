package main

import "fmt"

const NodesInBucket = 20

type Bucket struct {
	Contacts []*NodeId
}

func (b *Bucket) Describe() {
	for contact := range b.Contacts {
		fmt.Printf("NodeId %d: ", contact)
		b.Contacts[contact].Describe()
	}
}

func (b *Bucket) Add(contact *NodeId) {
	if len(b.Contacts) >= NodesInBucket {
		// if this happens we should actually ping each node and remove the slowest from the list instead of the last one
		b.Pop()
	}
	b.PushFront(contact)
}

// pop back item
func (b *Bucket) Pop() {
	b.Contacts = b.Contacts[1:]
}

func (b *Bucket) PushFront(contact *NodeId) {
	b.Contacts = append([]*NodeId{contact}, b.Contacts...)
}

func (b *Bucket) Has(id NodeId) (bool, int16) {
	if b.Contacts == nil {
		return false, -1
	}
	for i, contact := range b.Contacts {
		if contact.ID == id.ID {
			return true, int16(i)
		}
	}
	return false, -1
}

func (b *Bucket) Get(i int16) NodeId {
	return *b.Contacts[i]
}

func (b *Bucket) IsEmpty() bool {
	return len(b.Contacts) == 0
}
