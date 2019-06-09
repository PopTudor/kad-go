package main

import (
	"crypto/sha1"
	"fmt"
	"math/bits"
	"math/rand"
	"time"
)

const NodeLen = 20

type Id [NodeLen]byte

func NewNodeID() Id {
	var token [NodeLen]byte
	rand.Seed(time.Now().UnixNano())
	rand.Read(token[:])
	hasher := sha1.New()
	hasher.Write(token[:])
	sha := hasher.Sum(nil)
	copy(token[:], sha)
	return Id(token)
}
func NewNodeIdFrom(str string) Id {
	var token [NodeLen]byte
	copy(token[:], str)
	return Id(token)
}

// More shared bit pre-fix means closer distance between node ids
// This shared prefix will give leading zeros after the xor operation is done
//
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
func (nid *Id) SharedPrefixLen(oid Id) uint32 {
	prefix := 0
	for i := 0; i < NodeLen; i++ {
		xor := nid[i] ^ oid[i]
		leadingZeros := bits.LeadingZeros8(xor)
		prefix += leadingZeros
		//fmt.Printf("%08b %08b zeroes %08b\n", nid[i], oid[i], xor)
		if leadingZeros != 8 {
			//fmt.Println("break\n")
			break
		}
	}
	return uint32(prefix)
}
func (nid *Id) Describe() {
	fmt.Printf("NodeId: %s\n", nid.String())
}
func (nid *Id) DescribeHex() {
	fmt.Printf("%s\n", nid.StringHex())
}
func (nid *Id) StringHex() string {
	return fmt.Sprintf("%X", nid.Slice())
}
func (nid *Id) DescribeBinary() {
	fmt.Printf("%08b\n", nid.Slice())
}
func (nid *Id) Array() [NodeLen]byte {
	return [NodeLen]byte(*nid)
}
func (nid *Id) Slice() []byte {
	bytes := [NodeLen]byte(*nid)
	return bytes[:]
}
func (nid Id) String() string {
	return fmt.Sprintf("%X", [NodeLen]byte(nid))
}
