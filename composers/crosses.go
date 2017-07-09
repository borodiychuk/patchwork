package composers

import (
	"math/rand"

	"errors"

	m "github.com/borodiychuk/patchwork/models"
)

// Crosses composes a patter filled with crosses. Crosses guarantees that crosses do not have edges of similar pattern
type Crosses struct {
	seed         int64
	crosses      cross
	samplesCount int
	samples      []m.Sample
}

// Compose composes a picture out of provided set of samples
func (c *Crosses) Compose(canvas *m.Canvas, samples []m.Sample) error {
	// Prepare elements
	for i := 0; i < canvas.Width*canvas.Length; i++ {
		canvas.Elements = append(canvas.Elements, nil)
	}
	// Random point to start plotting
	x := rand.Intn(canvas.Width)
	y := rand.Intn(canvas.Length)
	// Initiate cross filling that will recursively call the same for neighbours
	cr := cross{
		center: coordinates{x: x, y: y},
		canvas: canvas,
	}
	if err := cr.fill(samples); err != nil {
		return err
	}
	return nil
}

// Seed seeds the ranom number generator
func (c *Crosses) Seed(seed int64) {
	rand.Seed(seed)
}

type coordinates struct {
	x int
	y int
}

type cross struct {
	center coordinates
	canvas *m.Canvas
}

func (c *cross) elements() [5]coordinates {
	return [5]coordinates{
		{x: c.center.x, y: c.center.y},
		{x: c.center.x - 1, y: c.center.y},
		{x: c.center.x + 1, y: c.center.y},
		{x: c.center.x, y: c.center.y - 1},
		{x: c.center.x, y: c.center.y + 1},
	}
}

func (c *cross) neighbours() [8]cross {
	return [8]cross{
		{canvas: c.canvas, center: coordinates{x: c.center.x - 1, y: c.center.y - 2}},
		{canvas: c.canvas, center: coordinates{x: c.center.x + 2, y: c.center.y - 1}},
		{canvas: c.canvas, center: coordinates{x: c.center.x + 1, y: c.center.y + 2}},
		{canvas: c.canvas, center: coordinates{x: c.center.x - 2, y: c.center.y + 1}},
		{canvas: c.canvas, center: coordinates{x: c.center.x + 1, y: c.center.y - 3}},
		{canvas: c.canvas, center: coordinates{x: c.center.x + 3, y: c.center.y + 1}},
		{canvas: c.canvas, center: coordinates{x: c.center.x - 1, y: c.center.y + 3}},
		{canvas: c.canvas, center: coordinates{x: c.center.x - 3, y: c.center.y + 1}},
	}
}

func (c *cross) fill(samples []m.Sample) error {
	// Find out forbidden samples for neighbout
	forbiddenSamples := [8]string{}
	for i, n := range c.neighbours() {
		forbiddenSamples[i] = n.sampleID()
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
		return errors.New("Seems too few samples to compose a nice pattern")
	}
	s := allowedSamples[rand.Intn(len(allowedSamples))]

	// Fill this cross with chosen sample
	outOfCanvas := 0
	for _, e := range c.elements() {
		// Do nothing if element is out of range
		i, err := c.canvas.IndexForXY(e.x, e.y)
		if err != nil {
			outOfCanvas++
			continue
		}
		// Do nothing if cross has already been processed. It is enough to declare that if at least one element is not nil
		if c.canvas.Elements[i] != nil {
			return nil
		}
		c.canvas.Elements[i] = &m.Element{
			Sample:   s,
			Rotation: rand.Intn(4),
		}
	}

	// Do not process neighbours if cross stays completely outside of canvas
	if outOfCanvas == 5 {
		return nil
	}

	// Process neighbours
	for _, n := range c.neighbours() {
		if err := n.fill(samples); err != nil {
			return err
		}
	}

	return nil
}

// sampleID returns an ID of sample that is used to fil the cross
func (c *cross) sampleID() string {
	for _, e := range c.elements() {
		i, err := c.canvas.IndexForXY(e.x, e.y)
		if err != nil {
			continue
		}
		if c.canvas.Elements[i] != nil && c.canvas.Elements[i].Sample != nil {
			return c.canvas.Elements[i].Sample.ID()
		}
	}
	return ""
}
