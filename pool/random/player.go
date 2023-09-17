package random

import (
	"github.com/herzrasen/common/player"
)

func (r *Random) Player() player.Player {
	return player.Player{
		Person: r.Person(),
		TechnicalAttrs: player.TechnicalAttrs{
			Catching:  r.Attr(),
			Dribbling: r.Attr(),
			Heading:   r.Attr(),
			Reflexes:  r.Attr(),
			Shooting:  r.Attr(),
			Throwing:  r.Attr(),
			Passing:   r.Attr(),
			Tackling:  r.Attr(),
			Technique: r.Attr(),
		},
		PhysicalAttrs: player.PhysicalAttrs{
			Acceleration:    r.Attr(),
			Agility:         r.Attr(),
			Balance:         r.Attr(),
			Jumping:         r.Attr(),
			InjuryProneness: r.Attr(),
			Pace:            r.Attr(),
			Stamina:         r.Attr(),
			Strength:        r.Attr(),
		},
		MentalAttrs: player.MentalAttrs{
			Aggression:    r.Attr(),
			Anticipation:  r.Attr(),
			Communication: r.Attr(),
			Composure:     r.Attr(),
			Decisions:     r.Attr(),
			Eccentricity:  r.Attr(),
			Focus:         r.Attr(),
			Leadership:    r.Attr(),
			Positioning:   r.Attr(),
		},
	}
}
