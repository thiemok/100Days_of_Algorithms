package eratosthenesSieve

import "math"

/**
 * Generates prime numbers up to the given n
 * @param		{uint}		n				Generates primes up to this number
 * @return	{[]uint}	primes	The generated primes
 */
func GeneratePrimes(n uint) (primes []uint) {

	primes = append(primes, 2)

	/**
	 * 1. Create list of odd integers from 3 trough n
	 */
	numbers := make([]bool, (n/2)-1)

	/**
	 * 2. let p equal 3
	 */
	/**
	 * 4. Find the first number greater than p in the list, that is not marked and use it as the new p.
	 * If there is no such number or p^2 is greater than n stop, otherwise repeat from step 3.
	 */
	sqrtN := uint(math.Sqrt(float64(n)))
	for i, j := uint(0), uint(3); j <= sqrtN; i, j = i+1, j+2 {
		if !numbers[i] {
			/**
			 * 3. Enumerate the multiples of p by counting from p^2 in incements of 2p
			 * and mark them in the list
			 */
			for k, l := (j*j/2)-1, j*j; l < n; k, l = k+j, l+2*j {
				numbers[k] = true
			}
		}
	}

	/**
	 * When the algorithm terminates, the remaining unmrked numbers in the list are all primes below n.
	 */
	for i, j := 0, uint(3); i < len(numbers); i, j = i+1, j+2 {
		if !numbers[i] {
			primes = append(primes, j)
		}
	}
	return primes
}
