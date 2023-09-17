package team

import (
	"encoding/json"
	"fmt"
	"github.com/herzrasen/common/player"
	"github.com/herzrasen/engine/point"
	"os"
)

type Position struct {
	BasePosition BasePosition  `json:"basePosition"`
	Location     point.Point   `json:"location"`
	Instructions []Instruction `json:"instructions"`
	PlayerId     string        `json:"playerId"`
}

type Card uint8

const (
	Yellow    Card = 0
	YellowRed Card = 1
	Red       Card = 2
)

type Booking struct {
	PlayerId string `json:"playerId"`
	Warning  bool   `json:"warning"`
	Card     *Card  `json:"card"`
}

type Substitution struct {
	SubbedOffPlayerId string `json:"subbedOffPlayerId"`
	BroughtOnPlayerId string `json:"broughtOnPlayerId"`
	Second            uint32 `json:"second"`
}

type Team struct {
	Id            string          `json:"id"`
	Positions     []Position      `json:"positions"`
	Players       []player.Player `json:"players"`
	Substitutes   []player.Player `json:"substitutes"`
	Substitutions []Substitution  `json:"substitutions"`
	Bookings      []Booking       `json:"bookings"`
	// substitutionRules, ...
}

func FromFile(teamFile string) (*Team, error) {
	data, err := os.ReadFile(teamFile)
	if err != nil {
		return nil, err
	}
	var team Team
	err = json.Unmarshal(data, &team)
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func (t *Team) PositionForPlayer(player *player.Player) (Position, error) {
	for _, pos := range t.Positions {
		if pos.PlayerId == player.Id {
			return pos, nil
		}
	}
	return Position{}, fmt.Errorf("no player with id %s found in team %s", player.Id, t.Id)
}
