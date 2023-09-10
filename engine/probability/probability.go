package probability

import "math/rand"

func Probability[E any](percentage int, onSuccess E, orElse E) E {
	p := float64(percentage) / 100
	if rand.Float64() < p {
		return onSuccess
	}
	return orElse
}
