package renderers

import (
	"image"
	"image/png"
	"log"
	"os"

	"github.com/borodiychuk/patchwork/models"
	"github.com/disintegration/imaging"
)

// PNG describes renderer that can output PNG data
type PNG struct {
	TargetFile string
}

// Render renders canvas into a PNG file
func (i *PNG) Render(canvas *models.Canvas) error {
	const patchSize = 50 // px
	var imageWidth, imageHeight = patchSize * canvas.Width, patchSize * canvas.Length

	output := image.NewNRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	// Iterate over elements
	for i := 0; i < canvas.Length*canvas.Width; i++ {
		e := canvas.Elements[i]
		col, row := canvas.XY(i)

		xOffset := col * patchSize
		yOffset := row * patchSize

		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		patchData := imaging.Resize(e.Sample.Image(), patchSize, patchSize, imaging.Lanczos)
		switch e.Rotation {
		case 1:
			patchData = imaging.Rotate90(patchData)
		case 2:
			patchData = imaging.Rotate180(patchData)
		case 3:
			patchData = imaging.Rotate270(patchData)
		}

		for y := 0; y < patchSize; y++ {
			for x := 0; x < patchSize; x++ {
				data := patchData.At(x, y)
				output.Set(xOffset+x, yOffset+y, data)
			}
		}

	}

	f, err := os.Create(i.TargetFile)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	if err := png.Encode(f, output); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
