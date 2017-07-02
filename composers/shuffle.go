package composers

import "github.com/borodiychuk/patchwork/models"

// Shuffler just picks random patch and turn it by random angle
type Shuffler struct {
}

// Compose composes a picture out of provided set of samples
func (c *Shuffler) Compose(*models.Canvas, []*models.Sample) {
	// Here goes the logic that fills canvas with patches
}

// Seed seends the ranom number generator
func (c *Shuffler) Seed(int64) {
}
