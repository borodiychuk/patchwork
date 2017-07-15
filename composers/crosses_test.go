package composers

import (
	"testing"

	m "github.com/borodiychuk/patchwork/models"
)

func TestCrosses_Compose(t *testing.T) {
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
			c := &Crosses{}
			samples := []m.Sample{}
			for i := 0; i < tt.count; i++ {
				samples = append(samples, &sample{})
			}
			if err := canvas.Compose(c, samples); (err != nil) != tt.wantErr {
				t.Errorf("Crosses.Compose() error = %v, wantErr %v", err, tt.wantErr)
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
