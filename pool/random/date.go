package random

import (
	"time"
)

func (r *Random) Date(minYearOfBirth int, maxYearOfBirth int) time.Time {
	year := r.IntInRange(minYearOfBirth, maxYearOfBirth)
	month := r.IntInRange(1, 12)
	day := r.IntInRange(1, 31)
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
