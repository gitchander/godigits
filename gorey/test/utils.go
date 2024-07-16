package main

import (
	"log"
	"math/rand"

	"github.com/gitchander/godigits/utils/random"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type IntNexter interface {
	NextInt() int
}

// ------------------------------------------------------------------------------
type serialIntNexter struct {
	a int
}

var _ IntNexter = &serialIntNexter{}

func (p *serialIntNexter) NextInt() int {
	b := p.a
	p.a++
	return b
}

func SerialIntNexterInit(min int) IntNexter {
	return &serialIntNexter{
		a: min,
	}
}

func SerialIntNexter() IntNexter {
	return SerialIntNexterInit(0)
}

// ------------------------------------------------------------------------------
type randomIntNexter struct {
	r *rand.Rand
}

var _ IntNexter = &randomIntNexter{}

func newRandomIntNexter() *randomIntNexter {
	return &randomIntNexter{
		r: random.NewRandNow(),
	}
}

func (p *randomIntNexter) NextInt() int {
	return p.r.Int()
}

func RandomIntNexter() IntNexter {
	return newRandomIntNexter()
}
