package interval

type IntervalValuer struct {
	min   float64
	delta float64
}

func NewIntervalValuer(l IntervalFloat, n int) *IntervalValuer {
	return &IntervalValuer{
		min:   l.Min,
		delta: l.Width() / float64(n-1),
	}
}

func (p *IntervalValuer) IndexToValue(i int) float64 {
	return p.min + (p.delta * float64(i))
}
