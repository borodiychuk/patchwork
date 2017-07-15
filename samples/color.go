package samples

import (
	"fmt"
	"image"
	"image/color"
)

// Color describes a sample that is just filled with particular color
type Color struct {
	Color color.Color
	image *image.RGBA
}

// Image generates 0px image of the given color
func (s *Color) Image() image.Image {
	if s.image == nil {
		s.image = image.NewRGBA(image.Rect(0, 0, 1, 1))
		s.image.Set(0, 0, s.Color)
	}
	return s.image
}

// ID returns identifier
func (s *Color) ID() string {
	r, g, b, a := s.Color.RGBA()
	return fmt.Sprintf("%d.%d.%d.%d", r, g, b, a)
}
