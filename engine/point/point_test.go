package point

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestPoint_DistanceTo(t *testing.T) {
	t.Run("distance to self should be 0", func(t *testing.T) {
		p := Point{
			X: 1,
			Y: 2,
		}
		distance := p.Distance(&p)
		assert.Equal(t, float64(0), distance)
	})

	t.Run("distance to other point", func(t *testing.T) {
		p1 := Point{
			X: 5,
			Y: 5,
		}
		p2 := Point{
			X: -5,
			Y: -5,
		}
		distance := p1.Distance(&p2)
		assert.Equal(t, math.Sqrt(10*10+10*10), distance)
	})
}
