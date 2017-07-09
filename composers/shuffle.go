package composers

import (
	"math/rand"

	"errors"

	m "github.com/borodiychuk/patchwork/models"
)

// Shuffle just picks random sample and turns it by random angle. Shuffle guarantees that similar elements have no common sides
type Shuffle struct{}

// Compose composes a picture out of provided set of samples
func (c *Shuffle) Compose(canvas *m.Canvas, samples []m.Sample) error {
	// Elements to exclude from list to avoid similar elements staying together
	forbiddenSamples := []string{}

	for i := 0; i < canvas.Length*canvas.Width; i++ {
		col, row, err := canvas.XYforIndex(i)
		if err != nil {
			return err
		}
		// This piece of code with cycle is not really optimal and could utilize some caching for better performance. But from other side,
		// that is not a tool for real-time data processing, so nothing bad happens if it stays so: not so fast
		// but semantically correct.
		allowedSamples := []m.Sample{}
		for _, s := range samples {
			allowed := true
			for _, fs := range forbiddenSamples {
				if s.ID() == fs {
					allowed = false
					break
				}
			}
			if allowed {
				allowedSamples = append(allowedSamples, s)
			}
		}
		if len(allowedSamples) == 0 {
			return errors.New("Too few samples to compose the pattern")
		}
		e := &m.Element{}
		e.Sample = allowedSamples[rand.Intn(len(allowedSamples))]
		e.Rotation = rand.Intn(4)
		canvas.Elements = append(canvas.Elements, e)

		forbiddenSamples = []string{e.Sample.ID()}
		if topElementIndex, err := canvas.IndexForXY(col, row-1); err == nil {
			forbiddenSamples = append(forbiddenSamples, canvas.Elements[topElementIndex].Sample.ID())
		}
		if topLeftElementIndex, err := canvas.IndexForXY(col-1, row-1); err == nil {
			forbiddenSamples = append(forbiddenSamples, canvas.Elements[topLeftElementIndex].Sample.ID())
		}
		if topRightElementIndex, err := canvas.IndexForXY(col+1, row-1); err == nil {
			forbiddenSamples = append(forbiddenSamples, canvas.Elements[topRightElementIndex].Sample.ID())
		}
	}
	return nil
}

// Seed seeds the ranom number generator
func (c *Shuffle) Seed(seed int64) {
	rand.Seed(seed)
}
