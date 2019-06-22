package main

import (
	"crypto/sha1"
	"fmt"
	"math/bits"
	"math/rand"
	"strconv"
	"time"
)

const KeyLen = 20

type Key [KeyLen]byte

func NewNodeKey() Key {
	var token [KeyLen]byte
	rand.Seed(time.Now().UnixNano())
	rand.Read(token[:])
	hasher := sha1.New()
	hasher.Write(token[:])
	sha := hasher.Sum(nil)
	copy(token[:], sha)
	return Key(token)
}
func NewKeyFrom(str string) Key {
	for i := len(str); i < KeyLen; i++ {
		str += "0"
	}
	r := stringToBin(str)
	var res [KeyLen]byte
	copy(res[:], r)
	return Key(res)
}

// Convert binary string to byte array
// eg "00000000000000000001" => 00000000000000000001
func stringToBin(s string) (binString []byte) {
	for _, c := range s {
		bin, _ := strconv.Atoi(string(c))
		binString = append(binString, byte(bin))
	}
	return
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
func (nid *Key) SharedPrefixLen(oid Key) uint32 {
	prefix := 0
	for i := 0; i < KeyLen; i++ {
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
func (nid *Key) Describe() {
	fmt.Printf("NodeId: %s\n", nid.String())
}
func (nid *Key) DescribeHex() {
	fmt.Printf("%s\n", nid.StringHex())
}
func (nid *Key) StringHex() string {
	return fmt.Sprintf("%X", nid.Slice())
}
func (nid *Key) DescribeBinary() {
	fmt.Printf("%08b\n", nid.Slice())
}
func (nid *Key) Array() [KeyLen]byte {
	return [KeyLen]byte(*nid)
}
func (nid *Key) Slice() []byte {
	bytes := [KeyLen]byte(*nid)
	return bytes[:]
}
func (nid Key) String() string {
	return fmt.Sprintf("%X", [KeyLen]byte(nid))
}
