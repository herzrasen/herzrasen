package random

import (
	"testing"
	"time"
)

func TestRandom_Date(t *testing.T) {
	r := Random{MinYear: 2000, MaxYear: 2010}
	AssertRandom(t, 1000, r.Date, func(date time.Time) bool {
		return date.After(time.Date(1999, 12, 31, 23, 59, 59, 0, time.UTC)) &&
			date.Before(time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC))
	})
}
