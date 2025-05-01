package gobits

import (
	"golang.org/x/exp/constraints"
)

type Unsigned = constraints.Unsigned

func SetBit[T Unsigned](x T, i int, b Bit) T {
	y := T(1) << i
	switch b {
	case 0:
		x &^= y
	case 1:
		x |= y
	default:
		panicBit(b)
	}
	return x
}

func GetBit[T Unsigned](x T, i int) Bit {
	return Bit((x >> i) & 1)
}

// //------------------------------------------------------------------------------

// // Uint8

// //------------------------------------------------------------------------------

// func Uint8SetBit(x uint8, i int, b Bit) uint8 {
// 	y := uint8(1) << i
// 	switch b {
// 	case 0:
// 		x &^= y
// 	case 1:
// 		x |= y
// 	default:
// 		panic(errInvalidBit(b))
// 	}
// 	return x
// }

// func Uint8GetBit(x uint8, i int) Bit {
// 	return Bit((x >> i) & 1)
// }

// //------------------------------------------------------------------------------

// // Uint16

// //------------------------------------------------------------------------------

// func Uint16SetBit(x uint16, i int, b Bit) uint16 {
// 	y := uint16(1) << i
// 	switch b {
// 	case 0:
// 		x &^= y
// 	case 1:
// 		x |= y
// 	default:
// 		panic(errInvalidBit(b))
// 	}
// 	return x
// }

// func Uint16GetBit(x uint16, i int) Bit {
// 	return Bit((x >> i) & 1)
// }

// //------------------------------------------------------------------------------

// // Uint32

// //------------------------------------------------------------------------------

// func Uint32SetBit(x uint32, i int, b Bit) uint32 {
// 	y := uint32(1) << i
// 	switch b {
// 	case 0:
// 		x &^= y
// 	case 1:
// 		x |= y
// 	default:
// 		panic(errInvalidBit(b))
// 	}
// 	return x
// }

// func Uint32GetBit(x uint32, i int) Bit {
// 	return Bit((x >> i) & 1)
// }

// //------------------------------------------------------------------------------

// // Uint64

// //------------------------------------------------------------------------------

// func Uint64SetBit(x uint64, i int, b Bit) uint64 {
// 	y := uint64(1) << i
// 	switch b {
// 	case 0:
// 		x &^= y
// 	case 1:
// 		x |= y
// 	default:
// 		panic(errInvalidBit(b))
// 	}
// 	return x
// }

// func Uint64GetBit(x uint64, i int) Bit {
// 	return Bit((x >> i) & 1)
// }

// //------------------------------------------------------------------------------
