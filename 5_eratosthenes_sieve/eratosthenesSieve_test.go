package eratosthenesSieve

import "testing"

func TestGeneratePrimes(t *testing.T) {
	values := []uint{15, 30, 256}
	results := [][]uint{
		[]uint{2, 3, 5, 7, 11, 13},
		[]uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29},
		[]uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
			31, 37, 41, 43, 47, 53, 59, 61, 67,
			71, 73, 79, 83, 89, 97, 101, 103, 107,
			109, 113, 127, 131, 137, 139, 149,
			151, 157, 163, 167, 173, 179, 181,
			191, 193, 197, 199, 211, 223, 227,
			229, 233, 239, 241, 251},
	}

	var primes []uint
	for i, v := range values {
		primes = GeneratePrimes(v)
		if !testEq(primes, results[i]) {
			t.Fatalf("Expected %v, but got %v", results[i], primes)
		}
	}

}

func testEq(a, b []uint) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
