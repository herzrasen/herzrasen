package match

import (
	"github.com/herzrasen/engine/probability"
	"github.com/herzrasen/engine/team"
)

type EventType uint8

const (
	CoinToss EventType = 0
)

type Event struct {
	Second           uint32
	EventType        EventType
	PossessingTeamId string
}

type Match struct {
	SecondsPassed uint32
	HomeTeam      *team.Team
	AwayTeam      *team.Team
	Events        []Event
	LastEvent     *Event
}

func NewMatch(homeTeam *team.Team, awayTeam *team.Team) *Match {
	return &Match{
		HomeTeam: homeTeam,
		AwayTeam: awayTeam,
	}
}

func (m *Match) Start() {
	coinTossWinner := m.coinToss()
	newEvent := Event{
		Second:           0,
		EventType:        CoinToss,
		PossessingTeamId: coinTossWinner.Id,
	}
	m.Events = append(m.Events, newEvent)
	m.LastEvent = &newEvent
}

func (m *Match) coinToss() *team.Team {
	return probability.Probability(50, m.HomeTeam, m.AwayTeam)
}
