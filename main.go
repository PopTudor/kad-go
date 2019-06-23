package main

import (
	"fmt"
	"math/bits"
)

func main() {
	n := bits.LeadingZeros8(1)
	fmt.Println(n)
}
