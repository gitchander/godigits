package gorey

func Triangle(object Object) Object {
	return RegularPolygon{
		N:       3,
		Phase:   0.25,
		Content: object,
	}
}
