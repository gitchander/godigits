package geom

import (
	"fmt"
	"strings"
)

func errAspectRatioValue(name string, v int) error {
	return fmt.Errorf("value %[1]s must be > 0: %[1]s = %[2]d", name, v)
}

func aspectRatioCheckValue(name string, v int) error {
	if v <= 0 {
		return errAspectRatioValue(name, v)
	}
	return nil
}

func errParseAspectRatio(s string) error {
	return fmt.Errorf("parse aspect ratio: invalid string (%s)", s)
}

//------------------------------------------------------------------------------

type AspectRatio struct {
	X int
	Y int
}

func MakeAspectRatio(x, y int) (AspectRatio, error) {

	if err := aspectRatioCheckValue("x", x); err != nil {
		return AspectRatio{}, err
	}
	if err := aspectRatioCheckValue("y", y); err != nil {
		return AspectRatio{}, err
	}

	g := gcd(x, y)
	if g != 0 {
		x /= g
		y /= g
	}
	a := AspectRatio{
		X: x,
		Y: y,
	}
	return a, nil
}

func MustAspectRatio(x, y int) AspectRatio {
	ar, err := MakeAspectRatio(x, y)
	if err != nil {
		panic(err)
	}
	return ar
}

func (a AspectRatio) Normalize() AspectRatio {
	return MustAspectRatio(a.X, a.Y)
}

//------------------------------------------------------------------------------

// func (a AspectRatio) CalcDx_(dy float64) (dx float64, ok bool) {
// 	if a.Y <= 0 {
// 		return 0, false
// 	}
// 	dx = dy * float64(a.X) / float64(a.Y)
// 	return dx, true
// }

// func (a AspectRatio) CalcDy_(dx float64) (dy float64, ok bool) {
// 	if a.X <= 0 {
// 		return 0, false
// 	}
// 	dy = dx * float64(a.Y) / float64(a.X)
// 	return dy, true
// }

func (a AspectRatio) MustCalcDx(dy float64) (dx float64) {
	if a.Y <= 0 {
		panic(errAspectRatioValue("y", a.Y))
	}
	dx = dy * float64(a.X) / float64(a.Y)
	return
}

func (a AspectRatio) MustCalcDy(dx float64) (dy float64) {
	if a.X <= 0 {
		panic(errAspectRatioValue("x", a.X))
	}
	dy = dx * float64(a.Y) / float64(a.X)
	return
}

//------------------------------------------------------------------------------

func (a AspectRatio) String() string {
	return fmt.Sprintf("%d:%d", a.X, a.Y)
}

func ParseAspectRatio(s string) (AspectRatio, error) {
	var zero AspectRatio
	vs := strings.Split(s, ":")
	if len(vs) != 2 {
		return zero, errParseAspectRatio(s)
	}
	x, err := parseInt(vs[0])
	if err != nil {
		return zero, errParseAspectRatio(s)
	}
	y, err := parseInt(vs[1])
	if err != nil {
		return zero, errParseAspectRatio(s)
	}
	return MakeAspectRatio(x, y)
}

func MustParseAspectRatio(s string) AspectRatio {
	a, err := ParseAspectRatio(s)
	if err != nil {
		panic(err)
	}
	return a
}

//------------------------------------------------------------------------------

type AspectRatioF struct {
	X float64
	Y float64
}

func MakeAspectRatioF(x, y float64) AspectRatioF {
	return AspectRatioF{
		X: x,
		Y: y,
	}
}

func (ar AspectRatioF) CalcDx(dy float64) (dx float64) {
	dx = dy * ar.X / ar.Y
	return
}

func (ar AspectRatioF) CalcDy(dx float64) (dy float64) {
	dy = dx * ar.Y / ar.X
	return
}
