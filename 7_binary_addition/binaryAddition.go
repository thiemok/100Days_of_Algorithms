package binaryAddition

func fullAdder(a bool, b bool, cIn bool) (s bool, cOut bool) {
	s = (a != b) != cIn
	cOut = (a && b) || (cIn && (a != b))
	return
}

/**
 * Performs binary addition of the given integers
 * @param		{uint}	a	The left operand
 * @param		{uint}	b	The right operand
 * @return	{uint}	The result of a + b
 */
func Add(a, b uint) uint {
	var res uint
	var c bool
	var s bool
	var i uint

	for (a != 0) || (b != 0) || c {
		s, c = fullAdder((a&1) == 1, (b&1) == 1, c)
		if s {
			res += (1 << i)
		}
		i++
		a >>= 1
		b >>= 1
	}

	return res
}
