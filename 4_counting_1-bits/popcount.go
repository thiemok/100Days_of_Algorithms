package popcount

func Popcount(value uint) (count uint) {
	for value > 0 {
		if (value & 1) == 1 {
			count++
		}
		value >>= 1
	}
	return count
}
