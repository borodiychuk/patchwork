package models

// Element describes a sample placed in a canvas
type Element struct {
	Rotation int // 0-3 means multiplication per 90 degree clockwise
	Sample   *Sample
}
