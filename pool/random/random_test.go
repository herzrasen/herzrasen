package random

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func AssertRandom[A any](t *testing.T, times uint32, randomFunc func() A, assertion func(a A) bool) {
	for i := 0; i < int(times); i++ {
		res := randomFunc()
		fmt.Println(res)
		assert.Truef(t, assertion(res), "assertion did not return true")
	}
}
