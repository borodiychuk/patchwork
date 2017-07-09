package models

import (
	"errors"
	"math"
)

// Canvas descrbes a patchwork that consists of samples and can be rendered.
// Here comes coordinates structure
//
//           X  0  1  2  |
//        Y  +-----------+----
//        0  |  0  1  2  | ^
//        1  |  3  4  5  | | LENGTH=4
//        2  |  6  7  8  | |
//        3  |  9 10 11  | V
//           +-----------+---
//           |< WIDTH=3 >|
//
type Canvas struct {
	Length   int        // First dimension
	Width    int        // Second dimension
	Elements []*Element // Canvas data
}

// Render renders canvas into some output using particular rendering logic
func (c *Canvas) Render(r Renderer) ([]byte, error) {
	if len(c.Elements) == 0 {
		return nil, errors.New("No canvas elements defined. Have you called the composer first?")
	}
	for _, e := range c.Elements {
		if e == nil {
			return nil, errors.New("Some canvas elements were not defined. Have you called the composer first?")
		}
	}
	return r.Render(c)
}

// Compose composes canvas elements out of samples based on particular composition logic
func (c *Canvas) Compose(cmp Composer, samples []Sample) error {
	return cmp.Compose(c, samples)
}

// XYforIndex returns element coordinates by its index. Opposite to GetIndex
func (c *Canvas) XYforIndex(i int) (int, int, error) {
	if i > c.Length*c.Width-1 || i < 0 {
		return -1, -1, errors.New("out of range")
	}
	row := int(math.Floor(float64(i) / float64(c.Width)))
	col := i - row*c.Width
	return col, row, nil
}

// IndexForXY returns patch index by its coordinates. Opposite to GetXY
func (c *Canvas) IndexForXY(x, y int) (int, error) {
	if x > c.Width-1 || y > c.Length-1 || x < 0 || y < 0 {
		return -1, errors.New("out of range")
	}
	return y*c.Width + x, nil
}
