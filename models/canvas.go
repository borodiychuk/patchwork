package models

// Canvas descrbe an actual canvas that can be rendered
type Canvas struct {
	Length   int       // First dimension
	Width    int       // Second dimension
	Elements []Element // Canvas data
}

// Render renders canvas ito some output using particular rendering logic
func (c *Canvas) Render(r Renderer) {
	r.Render(c)
}

// Compose composes canvas elements out of patches based on particular composition logic
func (c *Canvas) Compose(cmp Composer, samples []*Sample) {
	cmp.Compose(c, samples)
}
