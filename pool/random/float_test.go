package random

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloatInRange(t *testing.T) {
	r := Random{}
	for i := 0; i < 5000; i++ {
		f := r.FloatInRange(10, 20)
		assert.True(t, f >= 10, fmt.Sprintf("%f is not >= 10", f))
		assert.True(t, f <= 20, fmt.Sprintf("%f is not < 20", f))
	}
}
