package random

import "testing"

func TestRandom_Id(t *testing.T) {
	r := Random{}
	AssertRandom(t, 1000, r.Id, func(a string) bool {
		return a != ""
	})
}
