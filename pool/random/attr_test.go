package random

import "testing"

func TestRandom_Attr(t *testing.T) {
	r := Random{MinAttr: 1, MaxAttr: 5}
	AssertRandom(t, 1000, r.Attr, func(a float32) bool {
		return a >= 1 && a <= 5
	})
}
