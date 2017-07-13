package mcm

import "math"

/**
 * Structure representing the name and dimensions of a matrix
 * @property	{string}	name	The name of the matrix
 * @property	{uint}		rows	The number of rows
 * @property	{uint}		cols	The number of columns
 */
type MatrixDim struct {
	name string
	rows uint
	cols uint
}

/**
 * Structure representing the result of a matrix multiplication chain in terms of cost and dimensions
 * @property	{uint}	cost	The number of scalar multiplications needed to solve the chain
 * @property	{uint}	rows	The number of rows of the resulting matrix
 * @property	{uint}	cols	The number of columns of the resulting matrix
 */
type MultChain struct {
	cost uint
	rows uint
	cols uint
}

func (l MultChain) combine(r MultChain) MultChain {
	l.cost += r.cost + l.rows*l.cols*r.cols
	l.cols = r.cols
	return l
}

var defaultMultChain = MultChain{math.MaxUint64, 0, 0}

/**
 * Calculates the optimal grouping for multiplication of the given matrix chain.
 * And returns the optimal order of computation and its a {MultChain} struct representing its final cost and dimensions
 * @param		{[]MatrixDim}	chain					The chain for which the optimal grouping should be computed
 * @return	{string}			order					A string representing the optimal order of computation
 * @return	{MultChain}		optimalChain	The resulting cost and matrix size of the optimal order
 */
func MCM(chain []MatrixDim) (order string, optimalChain MultChain) {
	chainMemory := make(map[string]MultChain)

	return matrixChainOrder(chain, chainMemory)
}

func matrixChainToString(chain []MatrixDim) string {
	s := "("
	for _, value := range chain {
		s += value.name
	}
	s += ")"

	return s
}

func matrixChainOrder(chain []MatrixDim, chainMemory map[string]MultChain) (order string, optimalChain MultChain) {
	/**
	 * Check if the chain is only one element.
	 * If so return imediately with cost 0
	 */
	if len(chain) == 1 {
		matrix := chain[0]
		optimalChain = MultChain{0, matrix.rows, matrix.cols}
		chainMemory[matrix.name] = optimalChain
		return matrix.name, optimalChain
	}

	/**
	 * Run for each possible position at which the chain can be split
	 */
	var chainL, chainR []MatrixDim
	var orderL, orderR string
	var optimalChainL, optimalChainR MultChain
	results := make(map[string]MultChain)
	for i := 1; i < len(chain); i++ {
		chainL = chain[:i]
		chainR = chain[i:]

		/**
		 * Find minimum cost for each part of the chain
		 */
		var ok bool

		orderL = matrixChainToString(chainL)
		optimalChainL, ok = chainMemory[orderL]
		if !ok {
			orderL, optimalChainL = matrixChainOrder(chainL, chainMemory)
		}

		orderR = matrixChainToString(chainR)
		optimalChainR, ok = chainMemory[orderR]
		if !ok {
			orderR, optimalChainR = matrixChainOrder(chainR, chainMemory)
		}

		/**
		* Add the minimum costs and the cost for multiplying both parts together
		* and memorize the result
		 */
		order = "(" + orderL + orderR + ")"
		results[order] = optimalChainL.combine(optimalChainR)
		chainMemory[order] = results[order]
	}

	/**
	 * Find the result with the minimum cost and return it
	 */
	optimalChain = defaultMultChain
	for key, value := range results {
		if value.cost < optimalChain.cost {
			order = key
			optimalChain = value
		}
	}
	return order, optimalChain
}
