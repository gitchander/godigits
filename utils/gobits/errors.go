package gobits

import (
	"fmt"
)

func errInvalidBit(b Bit) error {
	return fmt.Errorf("invalid bit value %d", b)
}

func panicBit(b Bit) {
	panic(errInvalidBit(b))
}
