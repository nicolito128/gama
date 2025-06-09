package gama

import (
	"image/color"
	"slices"
)

type Bucket struct {
	arr []color.Color
}

func NewBucket() *Bucket {
	b := &Bucket{
		arr: make([]color.Color, 0),
	}
	return b
}

func (b *Bucket) Median() color.Color {
	if len(b.arr) == 0 {
		return color.RGBA{}
	}

	var reds, greens, blues, alphas []uint32
	for _, c := range b.arr {
		r, g, b, a := c.RGBA()
		reds = append(reds, r)
		greens = append(greens, g)
		blues = append(blues, b)
		alphas = append(alphas, a)
	}

	slices.Sort(reds)
	slices.Sort(greens)
	slices.Sort(blues)
	slices.Sort(alphas)

	medianR := getMedian(reds)
	medianG := getMedian(greens)
	medianB := getMedian(blues)
	medianA := getMedian(alphas)

	median := color.RGBA{
		R: uint8(medianR >> 8),
		G: uint8(medianG >> 8),
		B: uint8(medianB >> 8),
		A: uint8(medianA >> 8),
	}

	return median
}

func (b *Bucket) Push(c color.Color) {
	b.arr = append(b.arr, c)
}

func getMedian(s []uint32) uint32 {
	n := len(s)
	if n == 0 {
		return 0
	}
	if n%2 == 1 {
		return s[n/2]
	}
	return (s[n/2-1] + s[n/2]) / 2
}
