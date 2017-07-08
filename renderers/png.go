package renderers

import (
	"bytes"
	"errors"
	"image"
	"image/png"

	"github.com/borodiychuk/patchwork/models"
	"github.com/disintegration/imaging"
)

// PNG describes renderer that can output PNG data
type PNG struct{}

// Render renders canvas into a Image file
func (i *PNG) Render(canvas *models.Canvas) ([]byte, error) {
	const patchSize = 50 // px
	var imageWidth, imageHeight = patchSize * canvas.Width, patchSize * canvas.Length
	var col, row int
	var err error

	image := image.NewNRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	// Iterate over elements
	for i := 0; i < canvas.Length*canvas.Width; i++ {
		e := canvas.Elements[i]
		if col, row, err = canvas.XYforIndex(i); err != nil {
			return nil, err
		}

		xOffset := col * patchSize
		yOffset := row * patchSize

		patchData := imaging.Resize(e.Sample.Image(), patchSize, patchSize, imaging.Lanczos)
		switch e.Rotation {
		case 0:
		case 1:
			patchData = imaging.Rotate90(patchData)
		case 2:
			patchData = imaging.Rotate180(patchData)
		case 3:
			patchData = imaging.Rotate270(patchData)
		default:
			return nil, errors.New("Rotation is out of range")
		}

		for y := 0; y < patchSize; y++ {
			for x := 0; x < patchSize; x++ {
				data := patchData.At(x, y)
				image.Set(xOffset+x, yOffset+y, data)
			}
		}
	}

	buf := new(bytes.Buffer)
	if err = png.Encode(buf, image); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
