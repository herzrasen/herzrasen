package random

import "github.com/herzrasen/common/person"

func (r *Random) Height() person.Height {
	heightRandom := Random{minInt: r.MinHeight, maxInt: r.MaxHeight}
	return person.Height{
		Centimeters: uint32(heightRandom.Int()),
	}
}
