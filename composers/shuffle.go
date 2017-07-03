package composers

import "github.com/borodiychuk/patchwork/models"

// Shuffle just picks random sample and turns it by random angle
type Shuffle struct {
}

// Compose composes a picture out of provided set of samples
func (c *Shuffle) Compose(*models.Canvas, []models.Sample) {
	// Here goes the logic that fills canvas with patches
}

// Seed seends the ranom number generator
func (c *Shuffle) Seed(int64) {
}
