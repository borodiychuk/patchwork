package models

import "image"

// Sample defines an cloth sample that patchwork is made of
type Sample interface {
	Image() image.Image // Image data
	ID() string         // An identifier that allows to compare samples
}
