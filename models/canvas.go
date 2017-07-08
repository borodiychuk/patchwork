package models

import "math"

// Canvas descrbes a patchwork that consists of samples and can be rendered
type Canvas struct {
	Length   int        // First dimension
	Width    int        // Second dimension
	Elements []*Element // Canvas data
}

// Render renders canvas into some output using particular rendering logic
func (c *Canvas) Render(r Renderer) {
	r.Render(c)
}

// Compose composes canvas elements out of samples based on particular composition logic
func (c *Canvas) Compose(cmp Composer, samples []Sample) {
	cmp.Compose(c, samples)
}

// XY returns element coordinates by its index. Opposite to GetIndex
func (c *Canvas) XY(i int) (int, int) {
	row := int(math.Floor(float64(i) / float64(c.Width)))
	col := i - row*c.Width
	return col, row
}

// Index returns patch index by its coordinates. Opposite to GetXY
func (c *Canvas) Index(x, y int) int {
	return y*c.Width + x
}
