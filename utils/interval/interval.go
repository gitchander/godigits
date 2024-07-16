package interval

type IntervalFloat struct {
	Min, Max float64
}

func (v IntervalFloat) Width() float64 {
	return v.Max - v.Min
}
