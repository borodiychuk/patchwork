package renderers

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"log"
	"strings"
	"testing"

	m "github.com/borodiychuk/patchwork/models"
)

const base64PNG1x1 = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQMAAAAl21bKAAAABlBMVEUAAAD///+l2Z/dAAAACXBIWXMAAA7EAAAOxAGVKw4bAAAACklEQVQImWNgAAAAAgAB9HFkpgAAAABJRU5ErkJggg=="

type sample struct{}

func (s *sample) Image() image.Image {
	imageReader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64PNG1x1))
	imgData, _, err := image.Decode(imageReader)
	if err != nil {
		log.Fatal(err)
	}
	return imgData
}
func (s *sample) ID() string {
	return "sample"
}

func TestPNG_Render(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		width   int
		wantErr bool
	}{
		{name: "1x1", length: 1, width: 1},
		{name: "1x2", length: 1, width: 2},
		{name: "2x1", length: 2, width: 1},
		{name: "2x2", length: 2, width: 2},
		{name: "1x0", length: 1, width: 0, wantErr: true},
		{name: "0x1", length: 0, width: 1, wantErr: true},
		{name: "0x0", length: 0, width: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := PNG{}
			s := sample{}
			e := m.Element{Rotation: 1, Sample: &s}
			canvas := m.Canvas{
				Length:   tt.length,
				Width:    tt.width,
				Elements: []*m.Element{},
			}
			for i := 0; i < tt.length*tt.width; i++ {
				canvas.Elements = append(canvas.Elements, &e)
			}
			data, err := canvas.Render(&r)
			if (err != nil) != tt.wantErr {
				t.Errorf("PNG.Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			reader := bytes.NewReader(data)
			_, err = png.Decode(reader)
			if err != nil {
				t.Error("Data should be valid PNG image")
				return
			}

			reader = bytes.NewReader(data)
			cfg, _, err := image.DecodeConfig(reader)
			if err != nil {
				t.Error("Data should be have image config")
				return
			}
			if cfg.Width != tt.width*50 || cfg.Height != tt.length*50 {
				t.Error("wrong image dimensions")
			}
		})
	}
}
