package digits

import (
	"math/rand"

	"github.com/gitchander/godigits/utils/random"
)

func randNow() *rand.Rand {
	return random.NewRandNow()
}

func randBySeed(seed int64) *rand.Rand {
	return random.NewRandSeed(seed)
}

func randBool(r *rand.Rand) bool {
	return random.RandBool(r)
}

type testBase struct {
	Min int
	Max int
}

func randomBaseV1(r *rand.Rand) testBase {
	var (
		min = 0
		max = 20
	)
	b := testBase{
		Min: -random.RandIntIn(r, min, max),
		Max: +random.RandIntIn(r, min, max),
	}
	return b
}

func randomBase(r *rand.Rand) testBase {
	return randomBaseV1(r)
}
