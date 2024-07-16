package main

import (
	"sort"
)

func makeTableCountON() []int {
	var (
		bs = []byte("abcdefg")
		xs = make([]int, 256)
	)
	for i := range xs {
		x := uint8(i)
		var count int
		for _, b := range bs {
			if segmentIsOn(x, b) {
				count++
			}
		}
		xs[i] = count
	}
	return xs
}

func makeSortTable1() []uint8 {

	var (
		//bs = []byte("abcdefg")
		//bs = []byte("agdfbec")
		bs = []byte("feagdbc")
	)

	m := make(map[byte]int)
	for i, b := range bs {
		m[b] = i
	}

	xs := make([]uint8, 256)
	for i := range xs {
		x := uint8(i)
		var v uint8
		for b, offset := range m {
			if segmentIsOn(x, b) {
				v |= 1 << offset
			}
		}
		xs[i] = v
	}
	return xs
}

var (
	countTable = makeTableCountON()
	sortTable1 = makeSortTable1()
)

// ------------------------------------------------------------------------------
type sorter1 []uint8

func (x sorter1) Len() int { return len(x) }

func (x sorter1) Less(i, j int) bool {
	d := countTable[x[i]] - countTable[x[j]]
	if d < 0 {
		return true
	}
	if d == 0 {
		return sortTable1[x[i]] < sortTable1[x[j]]
	}
	return false
}

func (x sorter1) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func (x sorter1) Sort() { sort.Sort(x) }

// ------------------------------------------------------------------------------
func segmentsOr(x uint8, bs ...byte) bool {
	for _, b := range bs {
		if segmentIsOn(x, b) {
			return true
		}
	}
	return false
}

func segmentsAnd(x uint8, bs ...byte) bool {
	for _, b := range bs {
		if not(segmentIsOn(x, b)) {
			return false
		}
	}
	return true
}

func segmentsAreOn(x uint8, s string) int {
	var count int
	bs := []byte(s)
	for _, b := range bs {
		if segmentIsOn(x, b) {
			count++
		}
	}
	return count
}

// neighbours returns number of neighbours (b)
func neighbours(x uint8, b byte) int {
	switch b {
	case 'a':
		return segmentsAreOn(x, "bfg")
	case 'b':
		return segmentsAreOn(x, "acg")
	case 'c':
		return segmentsAreOn(x, "bdg")
	case 'd':
		return segmentsAreOn(x, "ce")
	case 'e':
		return segmentsAreOn(x, "dfg")
	case 'f':
		return segmentsAreOn(x, "aeg")
	case 'g':
		return segmentsAreOn(x, "bcef")
	default:
		return 0
	}
}

// ------------------------------------------------------------------------------
func generateDigits() []uint8 {
	var vs []uint8
	const n = 1 << 7
	for i := 0; i < n; i++ {
		v := uint8(i)
		if not(segmentsAnd(v, 'f', 'e')) {
			continue
		}
		if segmentsAnd(v, 'b', 'c') {
			continue
		}
		if not(segmentsOr(v, 'a', 'g', 'd')) {
			continue
		}
		if segmentIsOn(v, 'b') && (neighbours(v, 'b') == 0) {
			continue
		}
		if segmentIsOn(v, 'c') && (neighbours(v, 'c') == 0) {
			continue
		}

		vs = append(vs, v)
	}
	sorter1(vs).Sort()
	return vs
}
