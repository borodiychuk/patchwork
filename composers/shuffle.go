package composers

import (
	"math/rand"

	"github.com/borodiychuk/patchwork/models"
)

// Shuffle just picks random sample and turns it by random angle
type Shuffle struct {
	samplesCount int
	samples      []models.Sample
}

// Compose composes a picture out of provided set of samples
func (c *Shuffle) Compose(canvas *models.Canvas, samples []models.Sample) {
	// Here goes the logic that fills canvas with patches
	c.samplesCount = len(samples)
	c.samples = samples
	for i := 0; i < canvas.Length*canvas.Width; i++ {
		canvas.Elements = append(canvas.Elements, c.makeElement())
	}
}

// Seed seends the ranom number generator
func (c *Shuffle) Seed(seed int64) {
	rand.Seed(seed)
}

func (c *Shuffle) makeElement() *models.Element {
	e := &models.Element{}
	e.Sample = c.samples[rand.Intn(c.samplesCount)]
	e.Rotation = rand.Intn(4)
	return e
}
