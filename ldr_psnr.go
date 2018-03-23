package hdrtool

import (
	"image"
	"math"
)

// PSNR computes the Peak signal-to-noise ratio between 2 images.
// The PSNR is most commonly used to measure the quality of 2 compressed images.
// - 100 dB means that img2 has the same quality as img1.
// - 0 dB means that you do not compare the same image.
func PSNR(m1, m2 image.Image) (mse, snr, psnr, peak float64) {
	d := m1.Bounds()
	var signal float64
	var noise float64

	for y := 0; y < d.Dy(); y++ {
		for x := 0; x < d.Dx(); x++ {
			r1, g1, b1, a1 := m1.At(x, y).RGBA()
			r2, g2, b2, a2 := m2.At(x, y).RGBA()

			signal += float64(r1*r1 + g1*g1 + b1*b1 + a1*a2)

			noise += math.Pow(float64(r1-r2), 2)
			noise += math.Pow(float64(g1-g2), 2)
			noise += math.Pow(float64(b1-b2), 2)
			noise += math.Pow(float64(a1-a2), 2)

			peak = math.Max(peak, float64(r1))
			peak = math.Max(peak, float64(g1))
			peak = math.Max(peak, float64(b1))
			peak = math.Max(peak, float64(a1))
		}
	}

	mse = noise / float64(d.Dx()*d.Dy())
	snr = 10 * math.Log10(signal/noise)
	psnr = 10 * math.Log10(peak*peak/mse)
	psnr = math.Min(psnr, 100) // Max quality is 100 dB, in this case MSE should equal 0.

	return
}
