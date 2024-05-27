package geom

type Bounds struct {
	Min, Max Point
}

func MakeBounds(x1, y1, x2, y2 float64) Bounds {
	return Bounds{
		Min: MakePoint(x1, y1),
		Max: MakePoint(x2, y2),
	}
}

func (b Bounds) Dx() float64 {
	return b.Max.X - b.Min.X
}

func (b Bounds) Dy() float64 {
	return b.Max.Y - b.Min.Y
}

func (b Bounds) Size() Point {
	return Point{
		X: b.Max.X - b.Min.X,
		Y: b.Max.Y - b.Min.Y,
	}
}

func (b Bounds) Add(p Point) Bounds {
	return Bounds{
		Min: b.Min.Add(p),
		Max: b.Max.Add(p),
	}
}

func (b Bounds) Sub(p Point) Bounds {
	return Bounds{
		Min: b.Min.Sub(p),
		Max: b.Max.Sub(p),
	}
}
