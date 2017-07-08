package composers

import (
	"math/rand"

	"github.com/borodiychuk/patchwork/models"
)

// Shuffle just picks random sample and turns it by random angle. Shuffle guarantees that similar elements have no common sides
type Shuffle struct {
	samplesCount int
	samples      []models.Sample
}

// Compose composes a picture out of provided set of samples
func (c *Shuffle) Compose(canvas *models.Canvas, samples []models.Sample) error {
	// Here goes the logic that fills canvas with patches
	c.samplesCount = len(samples)
	c.samples = samples
	// Elements to avoid
	topElement := ""
	leftElement := ""
	for i := 0; i < canvas.Length*canvas.Width; i++ {
		col, row, err := canvas.XYforIndex(i)
		if err != nil {
			return err
		}
		if row > 0 {
			topElementIndex, err := canvas.IndexForXY(col, row-1)
			if err != nil {
				return err
			}
			topElement = canvas.Elements[topElementIndex].Sample.ID()
		}
		e := c.makeElement([2]string{topElement, leftElement})
		canvas.Elements = append(canvas.Elements, e)
		leftElement = e.Sample.ID()
	}
	return nil
}

// Seed seends the ranom number generator
func (c *Shuffle) Seed(seed int64) {
	rand.Seed(seed)
}

func (c *Shuffle) makeElement(except [2]string) *models.Element {
	e := &models.Element{}
	for {
		e.Sample = c.samples[rand.Intn(c.samplesCount)]
		if e.Sample.ID() != except[0] && e.Sample.ID() != except[1] {
			break
		}
	}
	e.Rotation = rand.Intn(4)
	return e
}
