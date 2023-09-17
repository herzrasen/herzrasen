package random

import "github.com/herzrasen/common/person"

func (r *Random) Weight() person.Weight {
	weightRandom := Random{minInt: r.MinWeight, maxInt: r.MaxWeight}
	return person.Weight{
		Grams: uint32(weightRandom.Int()),
	}
}
