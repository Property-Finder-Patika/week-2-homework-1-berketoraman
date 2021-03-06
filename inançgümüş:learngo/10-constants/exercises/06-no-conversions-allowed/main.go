package main

// EXERCISE: No Conversions Allowed
//  1. Fix the program without doing any conversion.
//  2. Explain why it doesn't work.

// EXPECTED OUTPUT
//  10h0m0s later...

import (
	"fmt"
	"time"
)

func main() {
	// hours's type is time.Duration
	// but later's type was int
	// now, `later` is typeless
	//
	// time.Duration's underlying type is int64
	// and, `later` is numeric typeless value
	// so, they can be multiplied
	const later = 10

	hours, _ := time.ParseDuration("1h")

	fmt.Printf("%s later...\n", hours*later)
}
