package geom

type Rectangle2f struct {
	Min Point2f
	Max Point2f
}

func PointToRect2f(p Point2f) Rectangle2f {
	return Rectangle2f{
		Min: p,
		Max: p,
	}
}

func (r Rectangle2f) Dx() float64 {
	return r.Max.X - r.Min.X
}

func (r Rectangle2f) Dy() float64 {
	return r.Max.Y - r.Min.Y
}

func (r Rectangle2f) Size() Point2f {
	return Point2f{
		X: r.Dx(),
		Y: r.Dy(),
	}
}

func (r Rectangle2f) Center() Point2f {
	return r.Min.Add(r.Max).DivScalar(2)
}

func (r Rectangle2f) Add(p Point2f) Rectangle2f {
	return Rectangle2f{
		Min: r.Min.Add(p),
		Max: r.Max.Add(p),
	}
}

func (r Rectangle2f) Shrink(f Frame) Rectangle2f {
	return Rectangle2f{
		Min: Point2f{
			X: r.Min.X + f.Left,
			Y: r.Min.Y + f.Top,
		},
		Max: Point2f{
			X: r.Max.X - f.Right,
			Y: r.Max.Y - f.Bottom,
		},
	}
}

func (r Rectangle2f) Grow(f Frame) Rectangle2f {
	return Rectangle2f{
		Min: Point2f{
			X: r.Min.X - f.Left,
			Y: r.Min.Y - f.Top,
		},
		Max: Point2f{
			X: r.Max.X + f.Right,
			Y: r.Max.Y + f.Bottom,
		},
	}
}
