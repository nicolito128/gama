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
	if len(b.arr) == 1 {
		return b.arr[0]
	}

	n := len(b.arr)
	reds := make([]uint8, 0, n)
	greens := make([]uint8, 0, n)
	blues := make([]uint8, 0, n)
	alphas := make([]uint8, 0, n)

	for _, c := range b.arr {
		r, g, b, a := c.RGBA()
		reds = append(reds, uint8(r>>8))
		greens = append(greens, uint8(g>>8))
		blues = append(blues, uint8(b>>8))
		alphas = append(alphas, uint8(a>>8))
	}

	slices.Sort(reds)
	slices.Sort(greens)
	slices.Sort(blues)
	slices.Sort(alphas)

	mid := len(b.arr) / 2
	if len(b.arr)%2 == 1 {
		return color.RGBA{reds[mid], greens[mid], blues[mid], alphas[mid]}
	}

	return color.RGBA{
		R: uint8((uint16(reds[mid-1]) + uint16(reds[mid])) / 2),
		G: uint8((uint16(greens[mid-1]) + uint16(greens[mid])) / 2),
		B: uint8((uint16(blues[mid-1]) + uint16(blues[mid])) / 2),
		A: uint8((uint16(alphas[mid-1]) + uint16(alphas[mid])) / 2),
	}
}

func (b *Bucket) Push(c color.Color) {
	b.arr = append(b.arr, c)
}
