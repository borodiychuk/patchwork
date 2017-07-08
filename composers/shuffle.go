package composers

import (
	"math/rand"

	"errors"

	"github.com/borodiychuk/patchwork/models"
)

// Shuffle just picks random sample and turns it by random angle. Shuffle guarantees that similar elements have no common sides
type Shuffle struct{}

// Compose composes a picture out of provided set of samples
func (c *Shuffle) Compose(canvas *models.Canvas, samples []models.Sample) error {
	// Elements to exclude from list to avoid similar elements staying together
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
		// This piece of code with cycle is not really optimal and could utilize some caching for better performance. But from other side,
		// that is not a tool for real-time data processing, so nothing bad happens if it stays so: not so fast
		// but semantically correct.
		allowedSamples := []models.Sample{}
		for _, s := range samples {
			if s.ID() != topElement && s.ID() != leftElement {
				allowedSamples = append(allowedSamples, s)
			}
		}
		if len(allowedSamples) == 0 {
			return errors.New("Too few samples to compose the pattern")
		}
		e := &models.Element{}
		e.Sample = allowedSamples[rand.Intn(len(allowedSamples))]
		e.Rotation = rand.Intn(4)
		canvas.Elements = append(canvas.Elements, e)
		leftElement = e.Sample.ID()
	}
	return nil
}

// Seed seends the ranom number generator
func (c *Shuffle) Seed(seed int64) {
	rand.Seed(seed)
}
