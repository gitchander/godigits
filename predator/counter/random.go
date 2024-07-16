package main

import (
	"math/rand"
	"time"
)

func NewRandNow() *rand.Rand {
	return NewRandTime(time.Now())
}

func NewRandTime(t time.Time) *rand.Rand {
	return NewRandSeed(t.UnixNano())
}

func NewRandSeed(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func RandBool(r *rand.Rand) bool {
	return (r.Uint32() & 1) == 1
}

func RandIntMinMax(r *rand.Rand, min, max int) int {
	return min + r.Intn(max-min)
}
