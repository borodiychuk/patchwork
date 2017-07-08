package composers

import (
	"encoding/base64"
	"image"
	"log"
	"strings"
	"testing"

	m "github.com/borodiychuk/patchwork/models"
	uuid "github.com/satori/go.uuid"
)

const base64PNG1x1 = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABAQMAAAAl21bKAAAABlBMVEUAAAD///+l2Z/dAAAACXBIWXMAAA7EAAAOxAGVKw4bAAAACklEQVQImWNgAAAAAgAB9HFkpgAAAABJRU5ErkJggg=="

type sample struct {
	id string
}

func (s *sample) Image() image.Image {
	imageReader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64PNG1x1))
	imgData, _, err := image.Decode(imageReader)
	if err != nil {
		log.Fatal(err)
	}
	return imgData
}
func (s *sample) ID() string {
	if s.id == "" {
		s.id = uuid.NewV4().String()
	}
	return s.id
}

func TestShuffle_Compose(t *testing.T) {
	tests := []struct {
		name    string
		count   int
		wantErr bool
	}{
		{name: "0", count: 0, wantErr: true},
		{name: "1", count: 1, wantErr: true},
		{name: "8", count: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			canvas := m.Canvas{
				Length: 4,
				Width:  3,
			}
			c := &Shuffle{}
			samples := []m.Sample{}
			for i := 0; i < tt.count; i++ {
				samples = append(samples, &sample{})
			}
			if err := canvas.Compose(c, samples); (err != nil) != tt.wantErr {
				t.Errorf("Shuffle.Compose() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if len(canvas.Elements) != 12 {
				t.Errorf("Canvas elements were not completely filled")
			}
		})
	}
}
