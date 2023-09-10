package gender

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromString(t *testing.T) {
	t.Run("male", func(t *testing.T) {
		g := FromString("male")
		assert.Equal(t, Male, g)
	})

	t.Run("female", func(t *testing.T) {
		g := FromString("female")
		assert.Equal(t, Female, g)
	})

	t.Run("undefined", func(t *testing.T) {
		g := FromString("hello")
		assert.Equal(t, Undefined, g)
	})
}
