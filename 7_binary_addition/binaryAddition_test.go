package binrayAddition

import "testing"

func TestAdd(t *testing.T) {
	aS := []uint{5, 12, 1337}
	bS := []uint{5, 3, 160}

	for i := 0; i < len(aS); i++ {
		a := aS[i]
		b := bS[i]
		res := Add(a, b)

		if res != uint(a+b) {
			t.Fatalf("Expected %d, but got %d for %d + %d", a+b, res, a, b)
		}
	}
}
