package random

import (
	"github.com/herzrasen/common/gender"
	"math/rand"
)

func (r *Random) FirstName() string {
	switch r.Gender {
	case gender.Male:
		idx := rand.Intn(len(r.NameConfig.FirstNames.Male))
		return r.NameConfig.FirstNames.Male[idx]
	case gender.Female:
		idx := rand.Intn(len(r.NameConfig.FirstNames.Female))
		return r.NameConfig.FirstNames.Female[idx]
	}
	return ""
}

func (r *Random) LastName() string {
	idx := rand.Intn(len(r.NameConfig.LastNames.Common))
	return r.NameConfig.LastNames.Common[idx]
}
