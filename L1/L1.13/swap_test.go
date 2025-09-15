package main

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestSwapMethodsAllPairs(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	edges := []int{0, -1, 1, math.MinInt64, math.MaxInt64}

	for _, aOrig := range edges {
		for _, bOrig := range edges {
			checkAllSwap(t, aOrig, bOrig)
		}
	}

	for i := 0; i < 100000000; i++ {
		aOrig := int(r.Int63() - r.Int63())
		bOrig := int(r.Int63() - r.Int63())
		checkAllSwap(t, aOrig, bOrig)
	}
}

func checkAllSwap(t *testing.T, aOrig, bOrig int) {
	a, b := swapAddSub(aOrig, bOrig)
	if a != bOrig || b != aOrig {
		t.Errorf("swapAddSub failed for a=%d, b=%d: got a=%d, b=%d", aOrig, bOrig, a, b)
	}

	a, b = swapXOR(aOrig, bOrig)
	if a != bOrig || b != aOrig {
		t.Errorf("swapXOR failed for a=%d, b=%d: got a=%d, b=%d", aOrig, bOrig, a, b)
	}

	// if aOrig != 0 && bOrig != 0 && aOrig != math.MinInt64 && aOrig != math.MaxInt64 && bOrig != math.MaxInt64 && bOrig != math.MinInt64 {
	// 	a, b = swapMulDiv(aOrig, bOrig)
	// 	if a != bOrig || b != aOrig {
	// 		t.Errorf("swapMulDiv failed for a=%d, b=%d: got a=%d, b=%d", aOrig, bOrig, a, b)
	// 	}
	// }
}
