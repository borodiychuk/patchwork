package models

import (
	"image"
)

// Canvas descrbe an actual canvas that can be rendered
type Canvas struct {
	Length int     // First dimension
	Width  int     // Second dimension
	Data   []Pixel // Canvas data
}

// Render renders canvas using particular renderer
func (c *Canvas) Render(r Renderer) (data []byte, contentType string) {
	return r.Render(c)
}

// Pixel describes a patch that the canvas consists of
type Pixel struct {
	Rotation int // 0-3 means multiplication per 90 degree clockwise
	Pattern  *Pattern
}

// Pattern defines an element that patchwork is built from
type Pattern interface {
	Image() *image.Image // Image data
	Amount() int         // Availabile amount of patches. -1 means infinite
}

// Renderer renders a canvas into specific format
type Renderer interface {
	Render(*Canvas) (data []byte, contentType string)
}

// Composer composes a canvas out of patches
type Composer interface {
	Compose(*Canvas, []Pattern)
}
