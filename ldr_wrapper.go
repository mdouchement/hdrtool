package hdrtool

import (
	"image"

	"github.com/mdouchement/hdr/hdrcolor"
)

// A LDRWrapper encapsulate an LDR image for hdr treatments.
type LDRWrapper struct {
	image.Image
}

// NewLDRWrapper instanciates a new LDRWrapper.
func NewLDRWrapper(m image.Image) *LDRWrapper {
	return &LDRWrapper{Image: m}
}

// Size returns the number of pixels
func (w *LDRWrapper) Size() int {
	return w.Bounds().Dx() * w.Bounds().Dy()
}

// HDRAt returns the pixel in LMS-space using CIE CAT02 matrix.
func (w *LDRWrapper) HDRAt(x, y int) hdrcolor.Color {
	r, g, b, _ := w.At(x, y).RGBA()
	return hdrcolor.RGB{R: float64(r) / 0xFFFF, G: float64(g) / 0xFFFF, B: float64(b) / 0xFFFF}
}
