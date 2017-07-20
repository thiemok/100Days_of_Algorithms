package hanoi

import (
	"math"
	"testing"

	"github.com/emirpasic/gods/stacks/arraystack"
)

var Ns = []uint{3, 7}

func TestHanoi(t *testing.T) {
	for _, n := range Ns {
		result, numSteps := Hanoi(n)
		testHanoiResult(t, n, result, numSteps)
	}
}

func testHanoiResult(t *testing.T, n uint, result [3]*arraystack.Stack, numSteps uint) {
	/**
	 * Check numSteps to be 2^n - 1
	 */
	expectedSteps := uint(math.Pow(2, float64(n))) - 1
	if numSteps != expectedSteps {
		t.Fatalf("Size %d failed with %d steps, but expected %d", n, numSteps, expectedSteps)
	}

	/**
	 * Check result to have n disks in ascending order on index 2
	 */
	if uint(result[2].Size()) == n {
		var lastSize uint
		it := result[2].Iterator()
		var value uint
		for it.Next() {
			value = it.Value().(uint)
			if value <= lastSize {
				t.Fatalf("Disks are in wrong order. Got size %d, but previous size was %d", value, lastSize)
			}
			lastSize = value
		}
	} else {
		t.Fatalf("Expected %d Disks a target, but got %d", n, result[2].Size())
	}
}
