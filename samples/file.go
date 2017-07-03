package samples

import (
	"image"
	"image/png"
	"log"
	"os"
)

// File describes sample that is stored in a file
type File struct {
	data image.Image
}

// Import reads file and sets respective variables
func (s *File) Import(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	s.data, err = png.Decode(file)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// Image returns a pointer to an image data
func (s *File) Image() image.Image {
	return s.data
}
