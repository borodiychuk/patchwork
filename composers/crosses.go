package composers

/*
import (
	"math/rand"

	"github.com/borodiychuk/patchwork/models"
)

// Crosses composes a patter filled with crosses. Crosses guarantees that crosses do not have edges of similar pattern
type Crosses struct {
	seed         int64
	crosses      cross
	samplesCount int
	samples      []models.Sample
}

// Compose composes a picture out of provided set of samples
func (c *Crosses) Compose(canvas *models.Canvas, samples []models.Sample) error {

	// find first cross randomly

	// Initiate cross filling that will recursively call the same for neighbours

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
func (c *Crosses) Seed(seed int64) {
	rand.Seed(seed)
}

func (c *Crosses) makeElement(except [2]string) *models.Element {
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

type cross struct {
	x         int
	y         int
	processed bool
}

func (c *cross) fill() {

}
*/
