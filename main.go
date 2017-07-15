package main

import (
	"flag"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	c "github.com/borodiychuk/patchwork/composers"
	m "github.com/borodiychuk/patchwork/models"
	r "github.com/borodiychuk/patchwork/renderers"
	s "github.com/borodiychuk/patchwork/samples"
	"github.com/borodiychuk/patchwork/support"
)

func main() {
	// Define flags and parse them
	var length, width int
	var seed int
	var output, composerType string
	var sampleFiles support.StringParamsSet
	var sampleColors support.StringParamsSet
	flag.IntVar(&length, "dim-l", 15, "Patches count per length dimension")
	flag.IntVar(&width, "dim-w", 12, "Patches count per widht dimension")
	flag.StringVar(&output, "out", "patchwork.png", "Path to output file")
	flag.StringVar(&composerType, "composer", "shuffle", "Pattern composer shuffle|crosses")
	flag.IntVar(&seed, "seed", 0, "Random seed to regenerate particular pattern. 0 or negative means seeding randomly")
	flag.Var(&sampleFiles, "sample-file", "Path to sample file. There can be multiple parameters of this type")
	flag.Var(&sampleColors, "sample-color", `Sample color definition in format of 3 ints 0-255 concatenated with comma, like  "0,1,255"`)
	flag.Parse()

	samples := []m.Sample{}

	// Process sample files
	for _, f := range sampleFiles {
		sample := &s.File{}
		err := sample.Import(f)
		if err != nil {
			log.Fatalln("! Unable to import from file:", f)
		}
		samples = append(samples, sample)
		log.Println("* Using sample file:", f)
	}

	// Process color-defined sample
	for _, c := range sampleColors {
		colors := strings.Split(c, ",")
		if len(colors) < 3 {
			log.Fatalln("! Wrong color definition:", c)
		}
		colorsInt := [3]uint8{}
		for i := 0; i < 3; i++ {
			// Base set to 9 in order to represent an extra bit given by uint8
			c, err := strconv.ParseInt(colors[i], 10, 9)
			if err != nil {
				log.Fatalln("! Can not parse color component:", colors[i])
			}
			colorsInt[i] = uint8(c)
		}
		sample := &s.Color{
			Color: color.RGBA{colorsInt[0], colorsInt[1], colorsInt[2], 255},
		}
		samples = append(samples, sample)
		log.Println("* Using sample color:", sample.ID())
	}

	// Prepare canvas composer. This is the one that composes the final look
	var composer m.Composer
	switch composerType {
	case "shuffle":
		composer = &c.Shuffle{}
	case "crosses":
		composer = &c.Crosses{}
	default:
		log.Fatalln("! Unknown composer:", composerType)
	}
	log.Println("* Using composer:", composerType)
	if seed < 1 {
		seed = time.Now().UTC().Nanosecond()
		log.Println("* Using random seed:", seed)
	}
	composer.Seed(int64(seed))

	// Prepare canvas rendered. This is the one that exports calvas into a readable format
	renderer := r.PNG{}

	// Generate canvas, copose pattern and render it
	canvas := m.Canvas{
		Length: length,
		Width:  width,
	}
	if err := canvas.Compose(composer, samples); err != nil {
		log.Fatalln("! Unable to compose pattern:", err)
	}
	data, err := canvas.Render(&renderer)
	if err != nil {
		log.Fatalln("! Unable to render result:", err)
	}

	// Dump the rendered image
	file, err := os.Create(output)
	if err != nil {
		log.Fatalln("! Unable to create result file:", err)
	}
	defer file.Close()

	if bytes, err := file.Write(data); err != nil {
		log.Fatalln("! Unable to write to the result file:", err)
	} else {
		log.Println("* Created target file with size:", bytes)
	}
}
