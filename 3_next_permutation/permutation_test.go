package permutation

import "testing"

func TestPermutations(t *testing.T) {
	tests := [][]string{
		[]string{"F", "A", "D", "E"},
		[]string{"F", "A", "E", "D"},
		[]string{"F", "D", "A", "E"},
	}
	results := [][]string{
		[]string{"F", "A", "E", "D"},
		[]string{"F", "D", "A", "E"},
		[]string{"F", "D", "E", "A"},
	}

	for i, v := range tests {
		testPermutation(v, results[i], t)
	}
}

func testPermutation(sequence []string, result []string, t *testing.T) {
	res := Permute(sequence)
	equal := true

	for i, v := range res {
		if v != result[i] {
			equal = false
			break
		}
	}

	if !equal {
		t.Fatalf("Expected %v, but got %v", result, res)
	}
}
