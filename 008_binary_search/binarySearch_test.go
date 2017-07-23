package binarySearch

import "testing"

func TestSearch(t *testing.T) {
	data := []uint{2, 3, 4, 8, 22, 23, 24, 25, 26, 28, 31, 39, 40, 43, 45, 49, 54, 58, 59, 60, 72, 73, 76, 87, 95, 97, 98}
	targets := []uint{22, 43, 98, 2, 1}
	results := []int{4, 13, 26, 0, -1}

	for i, target := range targets {
		res := Search(data, target)
		if res != results[i] {
			t.Fatalf("expected index %d for %d, but got %d", results[i], target, res)
		}
	}
}
