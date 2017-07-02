package patterns

import "image"

// File describes pattern that is stored in a file
type File struct {
	path   string // Path to pattern image
	amount int    // Amount of availabe patches of given pattern
}

// Import reads file and sets respective variables
func (p *File) Import(path string) error {
	// set path
	// set amount
	return nil
}

// Image returns a oinetr to an image data
func (p *File) Image() *image.Image {
	return nil
}

// Amount returns amount of available patches
func (p *File) Amount() int {
	return p.count
}
