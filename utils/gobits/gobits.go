package gobits

import (
	"fmt"
)

func errInvalidBit(b uint) error {
	return fmt.Errorf("invalid bit value %d", b)
}

//------------------------------------------------------------------------------

// Uint8

//------------------------------------------------------------------------------

func Uint8SetBit(x uint8, i int, b uint) uint8 {
	y := uint8(1) << i
	switch b {
	case 0:
		x &^= y
	case 1:
		x |= y
	default:
		panic(errInvalidBit(b))
	}
	return x
}

func Uint8GetBit(x uint8, i int) uint {
	return uint((x >> i) & 1)
}

//------------------------------------------------------------------------------

// Uint16

//------------------------------------------------------------------------------

func Uint16SetBit(x uint16, i int, b uint) uint16 {
	y := uint16(1) << i
	switch b {
	case 0:
		x &^= y
	case 1:
		x |= y
	default:
		panic(errInvalidBit(b))
	}
	return x
}

func Uint16GetBit(x uint16, i int) uint {
	return uint((x >> i) & 1)
}

//------------------------------------------------------------------------------

// Uint32

//------------------------------------------------------------------------------

func Uint32SetBit(x uint32, i int, b uint) uint32 {
	y := uint32(1) << i
	switch b {
	case 0:
		x &^= y
	case 1:
		x |= y
	default:
		panic(errInvalidBit(b))
	}
	return x
}

func Uint32GetBit(x uint32, i int) uint {
	return uint((x >> i) & 1)
}

//------------------------------------------------------------------------------

// Uint64

//------------------------------------------------------------------------------

func Uint64SetBit(x uint64, i int, b uint) uint64 {
	y := uint64(1) << i
	switch b {
	case 0:
		x &^= y
	case 1:
		x |= y
	default:
		panic(errInvalidBit(b))
	}
	return x
}

func Uint64GetBit(x uint64, i int) uint {
	return uint((x >> i) & 1)
}

//------------------------------------------------------------------------------
