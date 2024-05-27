package geom

type Point struct {
	X, Y float64
}

func MakePoint(x, y float64) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func (a Point) Add(b Point) Point {
	return Point{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a Point) Sub(b Point) Point {
	return Point{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}
