package models

// Composer composes a canvas out of samples by chosing them and placing in some order
type Composer interface {
	Compose(*Canvas, []*Sample)
}
