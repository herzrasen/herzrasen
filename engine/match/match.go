package match

import (
	"github.com/herzrasen/engine/event"
	"github.com/herzrasen/engine/probability"
	"github.com/herzrasen/engine/team"
)

type Match struct {
	Events []event.Event
}

func NewMatch(homeTeam *team.Team, awayTeam *team.Team) (*Match, error) {
	// initial positioning events
	var events []event.Event
	for _, player := range homeTeam.Players {
		position, err := homeTeam.PositionForPlayer(&player)
		if err != nil {
			return nil, err
		}
		events = append(events, event.Event{
			Second:    0,
			EventType: event.InitialPosition,
			Player:    &player,
			Team:      homeTeam,
			Location:  position.Location,
		})
	}
	for _, player := range awayTeam.Players {
		position, err := awayTeam.PositionForPlayer(&player)
		if err != nil {
			return nil, err
		}
		events = append(events, event.Event{
			Second:    0,
			EventType: event.InitialPosition,
			Player:    &player,
			Team:      awayTeam,
			Location:  position.Location,
		})
	}
	return &Match{
		Events: events,
	}, nil
}

func (m *Match) Start() {
	coinTossWinner := m.coinToss()
	m.Events = append(m.Events, event.Event{
		Second:    0,
		EventType: event.CoinToss,
		Team:      coinTossWinner,
	})
	// select players to perform kick off
	//kickOffPlayers := coinTossWinner.ActivePlayers.SelectClosestToLocation(m.BallLocation, 2)
	//
	//kp1 := &kickOffPlayers[0]
	//kp1.SetLocation(point.Point{Y: 1})
	//kp2 := &kickOffPlayers[1]
	//kp2.SetLocation(point.Point{Y: -1})
}

func (m *Match) coinToss() *team.Team {
	return probability.Probability(50, m.HomeTeam, m.AwayTeam)
}

func (m *Match) CurrentLocations(t *team.Team) []team.Position {

}
