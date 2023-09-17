package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFromBytes(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		data := []byte(`{
  "firstnames": {
    "male": [
      "Male"
    ],
    "female": [
      "Female"
    ]
  },
  "lastnames": {
    "common": [
      "Last"
    ]
  }
}
`)
		config, err := FromBytes(data)
		require.NoError(t, err)
		assert.Len(t, config.Firstnames.Male, 1)
		assert.Equal(t, "Male", config.Firstnames.Male[0])
		assert.Len(t, config.Firstnames.Female, 1)
		assert.Equal(t, "Female", config.Firstnames.Female[0])
		assert.Len(t, config.Lastnames.Common, 1)
		assert.Equal(t, "Last", config.Lastnames.Common[0])
	})
}
