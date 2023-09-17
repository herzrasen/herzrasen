package random

import (
	"github.com/herzrasen/common/person"
	"testing"
)

func TestRandom_Height(t *testing.T) {
	r := Random{MinHeight: 160, MaxHeight: 200}
	AssertRandom(t, 1000, r.Height, func(height person.Height) bool {
		return height.Centimeters >= 160 && height.Centimeters <= 200
	})
}
