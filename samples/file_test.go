package samples

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// File data samples
const (
	base64PNG1x1 = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQMAAAAl21bKAAAABlBMVEUAAAD///+l2Z/dAAAACXBIWXMAAA7EAAAOxAGVKw4bAAAACklEQVQImWNgAAAAAgAB9HFkpgAAAABJRU5ErkJggg=="
	base64NonPNG = "dGVzdA=="
)

var squarePNGFile, nonSquarePNGFile, notPNGFile *os.File

func TestMain(m *testing.M) {
	var err error

	if squarePNGFile, err = writeFileFromBase64(base64PNG1x1, "png_1"); err != nil {
		log.Fatal("Creating tempfile error:", err)
		return
	}
	defer os.Remove(squarePNGFile.Name())

	if notPNGFile, err = writeFileFromBase64(base64NonPNG, "png_2"); err != nil {
		log.Fatal("Creating tempfile error:", err)
		return
	}
	defer os.Remove(notPNGFile.Name())

	code := m.Run()
	os.Exit(code)
}

func TestFile_Import_GeneratesID(t *testing.T) {
	f1 := File{}
	f2 := File{}

	if err := f1.Import(squarePNGFile.Name()); err != nil {
		t.Error("Error importing file:", err)
	}
	if err := f2.Import(squarePNGFile.Name()); err != nil {
		t.Error("Error importing file:", err)
	}

	if f1.ID() == "" {
		t.Error("The sample ID is not generated")
	}
	if f1.ID() == f2.ID() {
		t.Error("IDs of different file instances musrt differ")
	}
}

func TestFile_Import_DoesntImportNotPNG(t *testing.T) {
	f := File{}
	if err := f.Import(notPNGFile.Name()); err == nil {
		t.Error("Import non-PNG file should produce an error")
	}
}

func TestFile_Import_DoesntImportNonexistFile(t *testing.T) {
	f := File{}
	if err := f.Import("/non/exist/file"); err == nil {
		t.Error("Import nonexisting file should produce an error")
	}
}

func TestFile_Import_ImportsProperFile(t *testing.T) {
	f := File{}
	if err := f.Import(squarePNGFile.Name()); err != nil {
		t.Error("Import proper file should not raise an error")
	}
	if f.Image() == nil {
		t.Error("Import must read image data")
	}
}

func writeFileFromBase64(data, prefix string) (*os.File, error) {
	imgData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	file, err := ioutil.TempFile("", prefix)
	if err != nil {
		return nil, err
	}
	if err := ioutil.WriteFile(file.Name(), imgData, 0644); err != nil {
		os.Remove(file.Name())
		return nil, err
	}
	return file, nil
}
