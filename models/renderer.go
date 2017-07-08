package models

// Renderer renders a canvas into specific format
type Renderer interface {
	Render(*Canvas) ([]byte, error)
}
