package monteCarloPi

import "testing"

func TestPi(t *testing.T) {
	tests := []uint{100, 1000, 100000000}

	for _, n := range tests {
		pi := Pi(n)
		t.Logf("got %f for %d iterations", pi, n)
	}
}
