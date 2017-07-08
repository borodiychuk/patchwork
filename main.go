package main

import (
	"flag"
	"log"
	"os"
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
	var output string
	var sampleFiles support.StringParamsSet
	flag.IntVar(&length, "dim-l", 15, "Patches count per length dimension")
	flag.IntVar(&width, "dim-w", 12, "Patches count per widht dimension")
	flag.StringVar(&output, "out", "patchwork.png", "Path to output file")
	flag.IntVar(&seed, "seed", 0, "Random seed to regenerate particular pattern. 0 or negative means seeding randomly")
	flag.Var(&sampleFiles, "sample", "Path to sample file. There can be multiple parameters of this type")
	flag.Parse()

	// TODO: Validate flags

	// Process samples
	samples := []m.Sample{}
	for i := 0; i < len(sampleFiles); i++ {
		sample := &s.File{}
		err := sample.Import(sampleFiles[i])
		if err != nil {
			log.Fatalln("! Unable to import from file:", sampleFiles[i])
		}
		samples = append(samples, sample)
		log.Println("* Using sample file:", sampleFiles[i])
	}

	// Prepare canvas composer. This is the one that composes the final look
	composer := c.Shuffle{}
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
	if err := canvas.Compose(&composer, samples); err != nil {
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
