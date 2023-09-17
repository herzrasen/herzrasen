package team

import (
	"github.com/herzrasen/engine/point"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestActivePlayers_SelectClosestToLocation(t *testing.T) {
	t.Run("select", func(t *testing.T) {
		var activePlayers ActivePlayers
		activePlayers = append(activePlayers, ActivePlayer{
			PlayerId: "a",
			Location: point.Point{
				X: 5,
				Y: 5,
			},
		})
		activePlayers = append(activePlayers, ActivePlayer{
			PlayerId: "b",
			Location: point.Point{
				X: 0,
				Y: 0,
			},
		})
		location := point.Point{
			X: -5,
			Y: -5,
		}
		closest := activePlayers.SelectClosestToLocation(location, 1)
		assert.Len(t, closest, 1)
		assert.Equal(t, "b", closest[0].PlayerId)
	})
}
