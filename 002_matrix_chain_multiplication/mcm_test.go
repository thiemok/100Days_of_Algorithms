package mcm

import "testing"

func TestMCM(t *testing.T) {
	var matrices = []MatrixDim{
		MatrixDim{"A", 10, 20},
		MatrixDim{"B", 20, 30},
		MatrixDim{"C", 30, 40},
	}
	_, optimalChain := MCM(matrices)

	/**
	 * Check cost to be 18000
	 */
	if optimalChain.cost != 18000 {
		t.Fatalf("Expected cost to be %d, bit got %d", 18000, optimalChain.cost)
	}

	/**
	 * Check resulting matrix dims to be 10 x 40
	 */
	if optimalChain.rows != 10 || optimalChain.cols != 40 {
		t.Fatalf("Excpected resulting matrix to be %d x %d, but got %d x %d", 10, 40, optimalChain.rows, optimalChain.cols)
	}
}
