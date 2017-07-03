package main

import (
	"flag"
	"time"

	"fmt"

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
			fmt.Println("! Unable to read file:", sampleFiles[i])
			return
		}
		samples = append(samples, sample)
		fmt.Println("* Using sample file:", sampleFiles[i])
	}

	// Prepare canvas composer. This is the one that composes the final look
	composer := c.Shuffle{}
	if seed < 1 {
		seed = time.Now().UTC().Nanosecond()
		fmt.Println("* Using random seed:", seed)
	}
	composer.Seed(int64(seed))

	// Prepare canvas rendered. This is the one that exports calvas into a readable format
	renderer := r.PNG{
		TargetFile: output,
	}

	// Generate canvas, copose pattern and render it
	canvas := m.Canvas{
		Length: length,
		Width:  width,
	}
	canvas.Compose(&composer, samples)
	canvas.Render(&renderer)
}
