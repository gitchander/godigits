package random

import (
	"crypto/rand"
	"encoding/binary"
)

var NextUint64 = func() func() uint64 {
	ch := make(chan uint64)
	go func() {
		const bytesPerUint64 = 8
		var (
			n    = 100
			data = make([]byte, (n * bytesPerUint64))
		)
		for {
			_, err := rand.Read(data)
			if err != nil {
				panic(err)
			}
			for i := 0; i < n; i++ {
				bs := data[(i * bytesPerUint64):]
				u := binary.BigEndian.Uint64(bs)
				ch <- u
			}
		}
	}()
	return func() uint64 {
		return <-ch
	}
}()

func NextSeed() int64 {
	u := NextUint64()
	u = u >> 1 // clear sign bit
	return int64(u)
}

func NextRand() *Rand {
	return NewRandSeed(NextSeed())
}
