package match

import (
	"github.com/herzrasen/engine/team"
)

type Action struct {
	Second           uint32
	Event            string
	PossessingTeamId string
}

type Match struct {
	SecondsPassed uint32
	HomeTeam      *team.Team
	AwayTeam      *team.Team
	Actions       []Action
}

func NewMatch(homeTeam *team.Team, awayTeam *team.Team) *Match {
	return &Match{
		HomeTeam: homeTeam,
		AwayTeam: awayTeam,
	}
}
