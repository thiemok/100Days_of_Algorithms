package popcount

import "testing"

func TestPopcount(t *testing.T) {
	values := []uint{60, 13, 12, 61, 49, 240, 15}
	results := []uint{4, 3, 2, 5, 3, 4, 4}

	var result, expectedResult uint
	for i, v := range values {
		result, expectedResult = Popcount(v), results[i]
		if result != expectedResult {
			t.Fatalf("Expected popcount of %d but got %d for %b", expectedResult, result, v)
		}
	}
}
