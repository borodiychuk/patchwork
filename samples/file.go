package samples

import (
	"image"
	"image/png"
	"os"

	uuid "github.com/satori/go.uuid"
)

// File describes sample that is stored in a file
type File struct {
	data image.Image
	id   string
}

// Import reads file and sets respective variables
func (s *File) Import(path string) error {
	var file *os.File
	var err error

	s.id = uuid.NewV4().String()

	if file, err = os.Open(path); err != nil {
		return err
	}
	defer file.Close()

	if s.data, err = png.Decode(file); err != nil {
		return err
	}

	return nil
}

// Image returns image data
func (s *File) Image() image.Image {
	return s.data
}

// ID returns identifier
func (s *File) ID() string {
	return s.id
}
