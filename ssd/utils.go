package main

import (
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func not(x bool) bool {
	return !x
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a % b
	return
}

func floorInt(x float64) int {
	return int(math.Floor(x))
}

func ceilInt(x float64) int {
	return int(math.Ceil(x))
}

func roundInt(x float64) int {
	return int(math.Round(x))
}

func randInts(n int, min, max int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := make([]int, n)
	for i := range a {
		a[i] = min + r.Intn(max-min)
	}
	return a
}

func RepeatByte(b byte, count int) []byte {
	bs := make([]byte, count)
	for i := range bs {
		bs[i] = b
	}
	return bs
}

func mod(a, b int) int {
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}

// ------------------------------------------------------------------------------
func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func formatInt(a int) string {
	return strconv.Itoa(a)
}

func parseFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
