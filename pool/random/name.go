package random

import (
	"github.com/herzrasen/common/gender"
	"math/rand"
)

func (r *Random) Firstname() string {
	switch r.Gender {
	case gender.Male:
		idx := rand.Intn(len(r.NameConfig.Firstnames.Male))
		return r.NameConfig.Firstnames.Male[idx]
	case gender.Female:
		idx := rand.Intn(len(r.NameConfig.Firstnames.Female))
		return r.NameConfig.Firstnames.Female[idx]
	}
	return ""
}

func (r *Random) Lastname() string {
	idx := rand.Intn(len(r.NameConfig.Lastnames.Common))
	return r.NameConfig.Lastnames.Common[idx]
}
