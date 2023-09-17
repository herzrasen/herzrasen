package match

import (
	"github.com/herzrasen/common/gender"
	"github.com/herzrasen/engine/event"
	"github.com/herzrasen/engine/point"
	"github.com/herzrasen/engine/team"
	"github.com/herzrasen/pool/config"
	"github.com/herzrasen/pool/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMatch_Start(t *testing.T) {
	t.Run("a coinToss event should be created", func(t *testing.T) {
		homeTeam := initTeam(t, true)
		awayTeam := initTeam(t, false)
		m, err := NewMatch(homeTeam, awayTeam)
		require.NoError(t, err)
		m.Start()
		assert.Equal(t, event.CoinToss, m.Events[0].EventType)
	})

}

func initTeam(t *testing.T, isHome bool) *team.Team {
	t.Helper()
	homeFactor := 1
	if isHome {
		homeFactor = -1
	}
	r := random.Random{
		Gender: gender.Male,
		NameConfig: &config.NameConfig{
			FirstNames: config.FirstNames{Male: []string{"A", "B", "C", "D", "E", "F"}},
			LastNames:  config.LastNames{Common: []string{"U", "V", "W", "X", "Y", "Z"}},
		},
	}
	var newTeam team.Team
	newTeam.Id = r.Id()
	for i := 0; i < 11; i++ {
		p := r.NewPlayer(0, 5, 2003, 2006)
		newTeam.Players = append(newTeam.Players, *p)
		if i == 0 {
			newTeam.Positions = append(newTeam.Positions, team.Position{
				BasePosition: team.GoalKeeper,
				Location: point.Point{
					X: int8(53 * homeFactor),
					Y: int8(0),
				},
				PlayerId: p.Id,
			})
		} else if i >= 1 && i < 5 {
			newTeam.Positions = append(newTeam.Positions, team.Position{
				BasePosition: team.Defender,
				Location: point.Point{
					X: int8(32 * homeFactor),
					Y: int8(-30 + (i * 20)),
				},
				PlayerId: p.Id,
			})
		} else if i >= 1 && i < 5 {
			newTeam.Positions = append(newTeam.Positions, team.Position{
				BasePosition: team.Defender,
				Location: point.Point{
					X: int8(32 * homeFactor),
					Y: int8(-30 + (i * 20)),
				},
				PlayerId: p.Id,
			})
		} else if i >= 5 && i < 9 {
			newTeam.Positions = append(newTeam.Positions, team.Position{
				BasePosition: team.Midfielder,
				Location: point.Point{
					X: int8(20 * homeFactor),
					Y: int8(-30 + (i * 20)),
				},
				PlayerId: p.Id,
			})
		} else if i >= 9 && i < 11 {
			newTeam.Positions = append(newTeam.Positions, team.Position{
				BasePosition: team.Striker,
				Location: point.Point{
					X: int8(10 * homeFactor),
					Y: int8(-10 + (i * 20)),
				},
				PlayerId: p.Id,
			})
		}
	}

	for i := 0; i < 5; i++ {
		p := r.NewPlayer(0, 5, 2003, 2006)
		newTeam.Substitutes = append(newTeam.Substitutes, *p)
	}

	return &newTeam
}
