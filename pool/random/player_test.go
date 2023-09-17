package random

import (
	"github.com/herzrasen/common/gender"
	"github.com/herzrasen/common/player"
	"github.com/herzrasen/pool/config"
	"testing"
)

func TestRandom_Player(t *testing.T) {
	r := Random{
		Gender: gender.Male,
		NameConfig: &config.NameConfig{
			Firstnames: config.Firstnames{
				Male: []string{"Foo"},
			},
			Lastnames: config.Lastnames{
				Common: []string{"Bar"},
			},
		},
		MinAttr:   1,
		MaxAttr:   5,
		MinYear:   2000,
		MaxYear:   2000,
		MinHeight: 180,
		MaxHeight: 180,
		MinWeight: 80000,
		MaxWeight: 80000,
	}
	AssertRandom(t, 1000, r.Player, func(p player.Player) bool {
		if p.Gender != gender.Male {
			return false
		}
		if p.Firstname != "Foo" {
			return false
		}
		if p.Lastname != "Bar" {
			return false
		}
		if p.Id == "" {
			return false
		}
		if p.Weight.Grams != 80000 {
			return false
		}
		if p.Height.Centimeters != 180 {
			return false
		}
		return true
	})
}
