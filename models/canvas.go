package models

// Canvas descrbes a patchwork that consists of samples and can be rendered
type Canvas struct {
	Length   int        // First dimension
	Width    int        // Second dimension
	Elements []*Element // Canvas data
}

// Render renders canvas into some output using particular rendering logic
func (c *Canvas) Render(r Renderer) {
	r.Render(c)
}

// Compose composes canvas elements out of samples based on particular composition logic
func (c *Canvas) Compose(cmp Composer, samples []Sample) {
	cmp.Compose(c, samples)
}
