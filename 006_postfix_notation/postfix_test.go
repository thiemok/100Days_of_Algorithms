package postfix

import "testing"

func TestSolve(t *testing.T) {
	terms := []string{
		"1 2 + 4 3 - + 10 5 / *",
		"1 2 * 6 2 / + 9 7 - ^",
		"1 2 3 4 5 + + + +",
	}
	results := []int{8, 25, 15}

	for i, v := range terms {
		res, e := Solve(v)
		if res != results[i] || e != nil {
			t.Fatalf("Expected %d,but got %d", results[i], res)
		}
	}
	_, e := Solve("12 3")
	if e == nil {
		t.Fatalf("Expected error on invalid input, but got none")
	}
}
