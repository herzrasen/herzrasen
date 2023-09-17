package random

import (
	"testing"
)

func TestFloatInRange(t *testing.T) {
	r := Random{minFloat: 1, maxFloat: 5}
	AssertRandom(t, 1000, r.Float, func(a float32) bool {
		return a >= 1 && a <= 5
	})
}
