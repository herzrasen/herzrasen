package random

import "github.com/herzrasen/common/player"

func (r *Random) NewPlayer(minAttrValue float32, maxAttrValue float32, minYearOfBirth int, maxYearOfBirth int) *player.Player {
	return &player.Player{
		Id:          r.Id(),
		FirstName:   r.FirstName(),
		LastName:    r.LastName(),
		Gender:      r.Gender,
		DateOfBirth: r.Date(minYearOfBirth, maxYearOfBirth),
		TechnicalAttrs: player.TechnicalAttrs{
			Catching:  r.FloatInRange(minAttrValue, maxAttrValue),
			Dribbling: r.FloatInRange(minAttrValue, maxAttrValue),
			Heading:   r.FloatInRange(minAttrValue, maxAttrValue),
			Reflexes:  r.FloatInRange(minAttrValue, maxAttrValue),
			Shooting:  r.FloatInRange(minAttrValue, maxAttrValue),
			Throwing:  r.FloatInRange(minAttrValue, maxAttrValue),
			Passing:   r.FloatInRange(minAttrValue, maxAttrValue),
			Tackling:  r.FloatInRange(minAttrValue, maxAttrValue),
			Technique: r.FloatInRange(minAttrValue, maxAttrValue),
		},
		PhysicalAttrs: player.PhysicalAttrs{
			Acceleration:    r.FloatInRange(minAttrValue, maxAttrValue),
			Agility:         r.FloatInRange(minAttrValue, maxAttrValue),
			Balance:         r.FloatInRange(minAttrValue, maxAttrValue),
			Jumping:         r.FloatInRange(minAttrValue, maxAttrValue),
			InjuryProneness: r.FloatInRange(minAttrValue, maxAttrValue),
			Pace:            r.FloatInRange(minAttrValue, maxAttrValue),
			Stamina:         r.FloatInRange(minAttrValue, maxAttrValue),
			Strength:        r.FloatInRange(minAttrValue, maxAttrValue),
		},
		MentalAttrs: player.MentalAttrs{
			Aggression:    r.FloatInRange(minAttrValue, maxAttrValue),
			Anticipation:  r.FloatInRange(minAttrValue, maxAttrValue),
			Communication: r.FloatInRange(minAttrValue, maxAttrValue),
			Composure:     r.FloatInRange(minAttrValue, maxAttrValue),
			Decisions:     r.FloatInRange(minAttrValue, maxAttrValue),
			Eccentricity:  r.FloatInRange(minAttrValue, maxAttrValue),
			Focus:         r.FloatInRange(minAttrValue, maxAttrValue),
			Leadership:    r.FloatInRange(minAttrValue, maxAttrValue),
			Positioning:   r.FloatInRange(minAttrValue, maxAttrValue),
		},
	}
}
