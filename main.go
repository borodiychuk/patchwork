package main

import (
	"flag"
	"time"

	"github.com/borodiychuk/patchwork/composers"
	"github.com/borodiychuk/patchwork/models"
	"github.com/borodiychuk/patchwork/renderers"
	"github.com/borodiychuk/patchwork/support"
)

func main() {
	// Define flags and parse them
	var length, width int
	var seed int64
	var output string
	var patterns support.StringParamsSet
	flag.IntVar(&length, "dim-l", 15, "Patches count per length dimension")
	flag.IntVar(&width, "dim-w", 12, "Patches count per widht dimension")
	flag.StringVar(&output, "out", "patchwork.png", "Path to output file")
	flag.Int64Var(&seed, "seed", 0, "Random seed to regenerate particular pattern. 0 or negative means seeding randomly")
	flag.Var(&patterns, "pattern", "Path to pattern file. There can be multiple parameters of this type")
	flag.Parse()

	// TODO: Validate flags

	// Chose samples
	// Get samples data

	// Prepare canvas composer. This is the one that composes the final look
	composer := composers.Shuffle{}
	if seed < 1 {
		seed = time.Now().UTC().UnixNano()
	}
	composer.Seed(seed)

	// Prepare canvas rendered. This si the one that exports calvas into a readable format
	renderer := renderers.PNG{}

	// Generate canvas, copose pattern and render it
	canvas := models.Canvas{}
	canvas.Compose(&composer, []*models.Sample{})
	canvas.Render(&renderer)
}
