package monteCarloPi

import (
	"math"
	"math/rand"
)

func Pi(n uint) float64 {
	var m uint
	for i := uint(0); i < n; i++ {
		d := math.Hypot(rand.Float64(), rand.Float64())
		if d < 1 {
			m++
		}
	}
	return 4 * float64(m) / float64(n)
}
