package main

type Node struct {
	Contact      *Contact
	RoutingTable *RoutingTable
	DHT          DHT
}

func NewNode() *Node {
	contact := NewContact()
	return &Node{
		Contact:      contact,
		RoutingTable: NewRoutingTable(contact),
	}
}
func NewNodeWithId(id NodeID) *Node {
	contact := NewContactWith(&id)
	return &Node{
		Contact:      contact,
		RoutingTable: NewRoutingTable(contact),
	}
}

// ping a node to find out if is online
func (n *Node) Ping(other *Node) {

}

// call to find a specific node with given id. The recipiend of this call
// looks in it's own routing table and returns a set of contacts that are closeset to
// the Contact that is being looked up
func (n *Node) FindNode(id NodeID) []Contact {
	return nil
}

// this call tries to find a specific file Contact to be located. If the receiving
// node finds this Contact in it's own DHT segment, it will return the corresponding
// URL. If not, the recipient node returns a list of contacts that are closest
// to the file Contact
func (n *Node) FindValue(value []byte) *FindValueResponse {
	return nil
}

// This call is used to store a key/value pair(fileID,location) in the DHT segment of the recipient node
// Upon each successful RPC, both the sending/receiving node insert/update each other's contact info in their
// own routing table
func (n *Node) Store(value FileID, contact Contact) {

}

/**
 *
 */
type FindValueResponse struct {
	ValueFound Segment
	Contacts   []Contact
}
