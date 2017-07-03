package samples

import "image"

// File describes sample that is stored in a file
type File struct {
	path string // Path to sample image
}

// Import reads file and sets respective variables
func (s *File) Import(path string) error {
	s.path = path
	return nil
}

// Image returns a pointer to an image data
func (s *File) Image() *image.Image {
	return nil
}
