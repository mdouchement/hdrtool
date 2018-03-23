package hdrtool

import (
	"image"
	"math"
)

var (
	// C1 covariance coeff for SSIM m1
	C1 = math.Pow(0.01, 2)
	// C2 covariance coeff for SSIM m2
	C2 = math.Pow(0.03, 2)
)

// SSIM computes the Structural SIMilarity Index between 2 images.
// The SSIM is most commonly used to measure the quality of 2 compressed images.
// - 1 means that img2 has the same quality as img1.
// - 0 means that you do not compare the same image.
func SSIM(m1, m2 image.Image) (ssim float64) {
	d := m1.Bounds()
	var s1 float64
	var s2 float64
	var ss float64
	var s12 float64

	for y := 0; y < d.Dy(); y++ {
		for x := 0; x < d.Dx(); x++ {
			r1, g1, b1, a1 := m1.At(x, y).RGBA()
			r2, g2, b2, a2 := m2.At(x, y).RGBA()

			sum1 := r1 + g1 + b1 + a1
			sum2 := r2 + g2 + b2 + a2

			s1 += float64(sum1)
			s2 += float64(sum2)

			ss += float64(sum1*sum1 + sum2*sum2)

			s12 += float64(sum1 * sum2)
		}
	}

	vari := ss - s1*s1 - s2*s2
	covar := s12 - s1*s2

	return (2*s1*s2 + C1) * (2*covar + C2) / ((s1*s1 + s2*s2 + C1) * (vari + C2))
}
