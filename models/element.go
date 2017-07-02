package models

// Element describes a patch that canvas consists of
type Element struct {
	Rotation int // 0-3 means multiplication per 90 degree clockwise
	Sample   *Sample
}
