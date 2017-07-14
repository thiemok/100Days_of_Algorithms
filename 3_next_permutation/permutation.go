package permutation

/**
 * Findes the next permutation of the given sequence
 * @param		{[]string}	values	The sequence for which the permutation should be found
 * @return	{[]string}	The next permutation
 */
func Permute(values []string) []string {
	n := len(values)
	pivot := 0

	/**
	 * Find the longest decreasing sequence, beginning at the end
	 */
	for i := 1; i < n; i++ {
		if values[n-i] >= values[(n-i)-1] {
			/**
			 * Denote preceding item as pivot
			 */
			pivot = (n - i) - 1
			break
		}
	}

	/**
	 * Swap pivot with the smallest higher item in the sequence
	 */
	seq := values[pivot+1:]
	for i := len(seq) - 1; i >= 0; i-- {
		if values[pivot] < seq[i] {
			values[pivot], seq[i] = seq[i], values[pivot]
			break
		}
	}

	/**
	 * Revert the sequence
	 */
	for i, j := 0, len(seq)-1; i < j; i, j = i+1, j-1 {
		seq[i], seq[j] = seq[j], seq[i]
	}

	return values
}
