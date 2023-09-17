package random

import (
	"time"
)

func (r *Random) Date() time.Time {
	yearRandom := Random{minInt: r.MinYear, maxInt: r.MaxYear}
	year := yearRandom.Int()
	monthRandom := Random{minInt: 1, maxInt: 12}
	month := monthRandom.Int()
	dayRandom := Random{minInt: 1, maxInt: 31}
	day := dayRandom.Int()
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
