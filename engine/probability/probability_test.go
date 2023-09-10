package probability

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestProbability(t *testing.T) {
	t.Run("50/50 chance", func(t *testing.T) {
		headsCount := 0
		tailsCount := 0
		for i := 0; i < 1000; i++ {
			res := Probability(50, "heads", "tails")
			if res == "heads" {
				headsCount++
			} else {
				tailsCount++
			}
		}
		fmt.Printf("heads: %d\n", headsCount)
		fmt.Printf("tails: %d\n", tailsCount)
		diff := math.Abs(float64(headsCount - tailsCount))
		fmt.Printf("diff: %f\n", diff)
		assert.True(t, diff < 100) // allow 10% deviation
	})

	t.Run("0% chance", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			res := Probability(0, false, true)
			if !res {
				t.FailNow()
			}
		}
	})

	t.Run("100% chance", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			res := Probability(100, true, false)
			if !res {
				t.FailNow()
			}
		}
	})

}
