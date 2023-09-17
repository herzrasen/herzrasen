package random

import "testing"

func TestRandom_IntInRange(t *testing.T) {
	r := Random{minInt: 1, maxInt: 100}
	AssertRandom(t, 1000, r.Int, func(a int) bool {
		return a >= 1 && a <= 100
	})
}
