package models

// Composer composes a canvas out of patches
type Composer interface {
	Compose(*Canvas, []*Sample)
}
