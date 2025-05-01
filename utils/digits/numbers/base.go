package numbers

// digits for interval: [min, max-1]
type Base struct {
	Min, Max int
}

func MakeBase(min, max int) Base {
	return Base{
		Min: min,
		Max: max,
	}
}

// Closed Max value: [min, max]
func BaseClosed(min, max int) Base {
	return MakeBase(min, (max + 1))
}

func (b Base) Empty() bool {
	return b.Min >= b.Max
}

func (b Base) Width() int {
	return b.Max - b.Min
}

var (
	Base2  = MakeBase(0, 2)
	Base10 = MakeBase(0, 10)

	// https://en.wikipedia.org/wiki/Balanced_ternary
	Bal3  = BaseClosed(-1, +1) // Balanced ternary
	Bal9  = BaseClosed(-4, +4)
	Bal27 = BaseClosed(-13, +13)
	Bal81 = BaseClosed(-40, +40)
)
