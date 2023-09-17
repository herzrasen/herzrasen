package event

import (
	"github.com/herzrasen/common/player"
	"github.com/herzrasen/engine/point"
	"github.com/herzrasen/engine/team"
)

type EventType uint8

const (
	InitialPosition EventType = 0
	CoinToss        EventType = 1
)

type Event struct {
	Second    uint32
	EventType EventType
	Player    *player.Player
	Team      *team.Team
	Location  point.Point
}
