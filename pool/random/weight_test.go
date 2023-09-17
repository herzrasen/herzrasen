package random

import (
	"github.com/herzrasen/common/person"
	"testing"
)

func TestRandom_Weight(t *testing.T) {
	r := Random{MinWeight: 6000, MaxWeight: 10000}
	AssertRandom(t, 1000, r.Weight, func(weight person.Weight) bool {
		return weight.Grams >= 6000 && weight.Grams <= 10000
	})
}
