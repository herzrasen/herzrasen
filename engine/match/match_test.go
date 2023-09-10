package match

import (
	"github.com/herzrasen/common/gender"
	"github.com/herzrasen/engine/team"
	"github.com/herzrasen/pool/config"
	"github.com/herzrasen/pool/random"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatch_Start(t *testing.T) {
	t.Run("a coinToss event should be created", func(t *testing.T) {
		homeTeam := initTeam(t)
		awayTeam := initTeam(t)
		m := Match{
			SecondsPassed: 0,
			HomeTeam:      homeTeam,
			AwayTeam:      awayTeam,
		}
		m.Start()
		assert.Equal(t, CoinToss, m.Events[0].EventType)
	})

}

func initTeam(t *testing.T) *team.Team {
	t.Helper()
	r := random.Random{
		Gender: gender.Male,
		NameConfig: &config.NameConfig{
			FirstNames: config.FirstNames{Male: []string{"A", "B", "C", "D", "E", "F"}},
			LastNames:  config.LastNames{Common: []string{"U", "V", "W", "X", "Y", "Z"}},
		},
	}
	var newTeam team.Team
	newTeam.Id = r.Id()
	for i := 0; i < 16; i++ {
		p := r.NewPlayer(0, 5, 2003, 2006)
		newTeam.Squad = append(newTeam.Squad, *p)
	}

	newTeam.ActivePlayers = append(newTeam.ActivePlayers, team.ActivePlayer{
		Position: team.Position{
			BasePosition: team.GoalKeeper,
		},
		PlayerId: newTeam.Squad[0].Id,
	})

	for i := 1; i < 5; i++ {
		newTeam.ActivePlayers = append(newTeam.ActivePlayers, team.ActivePlayer{
			Position: team.Position{
				BasePosition: team.Defender,
			},
			PlayerId: newTeam.Squad[i].Id,
		})
	}

	for i := 5; i < 9; i++ {
		newTeam.ActivePlayers = append(newTeam.ActivePlayers, team.ActivePlayer{
			Position: team.Position{
				BasePosition: team.Midfielder,
			},
			PlayerId: newTeam.Squad[i].Id,
		})
	}

	for i := 9; i < 11; i++ {
		newTeam.ActivePlayers = append(newTeam.ActivePlayers, team.ActivePlayer{
			Position: team.Position{
				BasePosition: team.Striker,
			},
			PlayerId: newTeam.Squad[i].Id,
		})
	}
	return &newTeam
}
