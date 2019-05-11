package main

import (
	"crypto/sha1"
	"fmt"
	"math/bits"
	"math/rand"
	"time"
)

const NodeLen = 20

type NodeID [NodeLen]byte

func NewNodeID() NodeID {
	var token [NodeLen]byte
	var token2 [NodeLen]byte
	rand.Seed(time.Now().UnixNano())
	rand.Read(token[:])
	hasher := sha1.New()
	hasher.Write(token[:])
	sha := hasher.Sum(nil)
	copy(token2[:], sha)
	return NodeID(token2)
}

// More shared bit pre-fix means closer distance between node ids
// This shared prefix will give leading zeros after the xor operation is done
// 0 xor 0 = 0
// 0 xor 1 = 1
// 1 xor 0 = 1
// 1 xor 1 = 0
// ex: D(11,10)
// 11: 1011
// 10: 1010
// xor---------
//     0001 = 1 distance
// we have 3 shared bits and distance is only 1! if we were to have less shared bits
// the distance would have been greater
// Notice that the longer the shared sequence of bits is, the more zeroes we have
// in the resulting number
func (nid *NodeID) SharedPrefixLen(oid *NodeID) uint32 {
	var prefix uint32 = 0
	for i := 0; i < NodeLen; i++ {
		xor := nid[i] ^ oid[i]
		leadingZeros := bits.LeadingZeros8(xor)
		prefix += uint32(leadingZeros)
		//fmt.Printf("%08b %08b zeroes %08b\n", nid[i], oid[i], xor)
		if leadingZeros == 0 {
			//fmt.Println("break\n")
			break
		}
	}
	return prefix
}
func (nid *NodeID) Describe() {
	fmt.Printf("%s\n", nid.String())
}
func (nid *NodeID) DescribeHex() {
	fmt.Printf("%X\n", nid.Slice())
}
func (nid *NodeID) DescribeBinary() {
	fmt.Printf("%08b\n", nid.Slice())
}
func (nid *NodeID) Array() [NodeLen]byte {
	return [NodeLen]byte(*nid)
}
func (nid *NodeID) Slice() []byte {
	bytes := [NodeLen]byte(*nid)
	return bytes[:]
}
func (nid *NodeID) String() string {
	return fmt.Sprintf("%v", [NodeLen]byte(*nid))
}
