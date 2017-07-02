package samples

import "image"

// File describes sample that is stored in a file
type File struct {
	path   string // Path to sample image
	amount int    // Amount of availabe patches of given sample
}

// Import reads file and sets respective variables
func (s *File) Import(path string) error {
	// set path
	// set amount
	return nil
}

// Image returns a oinetr to an image data
func (s *File) Image() *image.Image {
	return nil
}

// Amount returns amount of available patches
func (s *File) Amount() int {
	return s.count
}
