package binarySearch

/**
 * Searches for the target in the given sorted array and returns its index or -1 if not found
 * @param		{[]uint}	data		The data to be searched in
 * @paran		{uint}		target	The target to be searched for
 * @return	{int}			index		The index of target in data or -1 if not found
 */
func Search(data []uint, target uint) (index int) {
	var l, m int
	var r = len(data) - 1
	var aM uint

	for {
		if l > r {
			return -1
		}

		m = (l + r) / 2
		aM = data[m]

		if aM == target {
			return m
		}

		if aM < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
}
