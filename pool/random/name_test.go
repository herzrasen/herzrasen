package random

import (
	"github.com/herzrasen/common/gender"
	"github.com/herzrasen/pool/config"
	"testing"
)

func TestRandom_Firstname(t *testing.T) {
	t.Run("male", func(t *testing.T) {
		r := Random{NameConfig: &config.NameConfig{
			Firstnames: config.Firstnames{Male: []string{"Albert", "Franz", "Xaver"}},
		}}
		AssertRandom(t, 1000, r.Firstname, func(name string) bool {
			for _, n := range r.NameConfig.Firstnames.Male {
				if n == name {
					return true
				}
			}
			return false
		})
	})

	t.Run("female", func(t *testing.T) {
		r := Random{Gender: gender.Female,
			NameConfig: &config.NameConfig{
				Firstnames: config.Firstnames{
					Female: []string{"Laura", "Elisabeth"},
				},
			}}
		AssertRandom(t, 1000, r.Firstname, func(name string) bool {
			for _, n := range r.NameConfig.Firstnames.Female {
				if n == name {
					return true
				}
			}
			return false
		})
	})

	t.Run("empty", func(t *testing.T) {
		r := Random{Gender: gender.Undefined,
			NameConfig: &config.NameConfig{
				Firstnames: config.Firstnames{
					Male:   []string{"Albert", "James"},
					Female: []string{"Laura", "Elisabeth"},
				},
			}}
		AssertRandom(t, 1000, r.Firstname, func(name string) bool {
			return name == ""
		})
	})
}

func TestRandom_Lastname(t *testing.T) {
	r := Random{NameConfig: &config.NameConfig{
		Firstnames: config.Firstnames{Male: []string{}},
		Lastnames:  config.Lastnames{Common: []string{"Smith", "Brown", "Nguyen", "Müller"}},
	}}
	AssertRandom(t, 1000, r.Lastname, func(name string) bool {
		for _, n := range r.NameConfig.Lastnames.Common {
			if n == name {
				return true
			}
		}
		return false
	})
}
