package renderers

import "github.com/borodiychuk/patchwork/models"

// PNG describes renderer that can output PNG data
type PNG struct{}

// Render renders canvas into a PNG file
func (i *PNG) Render(c *models.Canvas) {

}
