package random

import (
	"math/rand"
)

type Rand = rand.Rand

func RandBool(r *Rand) bool {
	return (r.Uint32() & 1) == 1
}
