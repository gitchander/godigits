package colorf

import (
	"fmt"
	"image/color"
	"math"
)

const colorComponentMax = 0xffff

type Colorf struct {
	R, G, B, A float64
}

var _ color.Color = Colorf{}

func (c Colorf) RGBA() (r, g, b, a uint32) {
	u := color.NRGBA64{
		R: colorComponentConvert(c.R),
		G: colorComponentConvert(c.G),
		B: colorComponentConvert(c.B),
		A: colorComponentConvert(c.A),
	}
	return u.RGBA()
}

func colorComponentConvert(v float64) uint16 {
	return uint16(math.Round(clampFloat64(v, 0, 1) * colorComponentMax))
}

func MakeColorf(r, g, b, a float64) color.Color {
	return Colorf{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

func ParseColor(s string) (color.Color, error) {

	bs := []byte(s)
	if firstByteIsNot(bs, '#') {
		return nil, fmt.Errorf("invalid color (%s): no symbol %q", s, '#')
	}
	bs = bs[1:]

	ns, err := decodeNibbles(bs)
	if err != nil {
		return nil, fmt.Errorf("invalid color (%s): %v", s, err)
	}

	var c color.Color
	switch k := len(ns); k {
	case 3: // rgb
		c = color.NRGBA{
			R: nibblesToByte(ns[0], ns[0]),
			G: nibblesToByte(ns[1], ns[1]),
			B: nibblesToByte(ns[2], ns[2]),
			A: 0xff,
		}
	case 4: // rgba
		c = color.NRGBA{
			R: nibblesToByte(ns[0], ns[0]),
			G: nibblesToByte(ns[1], ns[1]),
			B: nibblesToByte(ns[2], ns[2]),
			A: nibblesToByte(ns[3], ns[3]),
		}
	case 6: // rrggbb
		c = color.NRGBA{
			R: nibblesToByte(ns[0], ns[1]),
			G: nibblesToByte(ns[2], ns[3]),
			B: nibblesToByte(ns[4], ns[5]),
			A: 0xff,
		}
	case 8: // rrggbbaa
		c = color.NRGBA{
			R: nibblesToByte(ns[0], ns[1]),
			G: nibblesToByte(ns[2], ns[3]),
			B: nibblesToByte(ns[4], ns[5]),
			A: nibblesToByte(ns[6], ns[7]),
		}
	default:
		return nil, fmt.Errorf("invalid color (%s): invalid number of nibbles", s)
	}
	return c, nil
}

func MustParseColor(s string) color.Color {
	c, err := ParseColor(s)
	if err != nil {
		panic(err)
	}
	return c
}

func decodeNibble(b byte) (byte, bool) {
	if ('0' <= b) && (b <= '9') {
		return b - '0', true
	}
	if ('a' <= b) && (b <= 'f') {
		return b - 'a' + 10, true
	}
	if ('A' <= b) && (b <= 'F') {
		return b - 'A' + 10, true
	}
	return 0, false
}

func decodeNibbles(bs []byte) ([]byte, error) {
	ns := make([]byte, len(bs))
	for i, b := range bs {
		n, ok := decodeNibble(b)
		if !ok {
			return nil, fmt.Errorf("invalid nibble %#U", b)
		}
		ns[i] = n
	}
	return ns, nil
}

func firstByteIsNot(bs []byte, b byte) bool {
	return not(firstByteIs(bs, b))
}

func firstByteIs(bs []byte, b byte) bool {
	return (len(bs) > 0) && (bs[0] == b)
}
