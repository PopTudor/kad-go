package main

import "fmt"

const NodesInBucket = 20

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
		b.Pop()
	}
	b.PushFront(contact)
}

// pop back item
func (b *Bucket) Pop() {
	b.Contacts = b.Contacts[1:]
}

func (b *Bucket) PushFront(contact *Contact) {
	b.Contacts = append([]*Contact{contact}, b.Contacts...)
}
