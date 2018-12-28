package cmd

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"time"

	// Import LDR codecs
	_ "image/jpeg"
	_ "image/png"

	"github.com/mdouchement/hdr"
	"github.com/mdouchement/hdrtool"

	// Import HDR codecs
	_ "github.com/mdouchement/hdr/codec/crad"
	_ "github.com/mdouchement/hdr/codec/pfm"
	_ "github.com/mdouchement/hdr/codec/rgbe"
	_ "github.com/mdouchement/tiff"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	// QualityCommand defines the command for compute quality check.
	QualityCommand = &cobra.Command{
		Use:   "quality [flags] source_file target_file",
		Short: "Performs PSNR and SSIM",
		Long:  "Performs PSNR and SSIM",
		RunE:  qualityAction,
	}
)

func qualityAction(c *cobra.Command, args []string) error {
	if len(args) != 2 {
		return errors.New("quality: Invalid number of arguments")
	}
	f1, err := os.Open(args[0])
	if err != nil {
		return errors.Wrap(err, "quality:")
	}
	defer f1.Close()

	start := time.Now()
	m1, fname, err := image.Decode(f1)
	if err != nil {
		return errors.Wrap(err, "quality:")
	}
	fmt.Printf("Read image (%dx%dp - %s - %v) %s\n", m1.Bounds().Dx(), m1.Bounds().Dy(), fname, time.Since(start), filepath.Base(args[0]))

	f2, err := os.Open(args[1])
	if err != nil {
		return errors.Wrap(err, "quality:")
	}
	defer f2.Close()

	start = time.Now()
	m2, fname, err := image.Decode(f2)
	if err != nil {
		return errors.Wrap(err, "quality:")
	}
	fmt.Printf("Read image (%dx%dp - %s - %v) %s\n", m2.Bounds().Dx(), m2.Bounds().Dy(), fname, time.Since(start), filepath.Base(args[1]))

	var mse, snr, psnr, peak, ssim float64
	hm1, ok1 := m1.(hdr.Image)
	hm2, ok2 := m2.(hdr.Image)
	if ok1 && ok2 {
		mse, snr, psnr, peak = hdrtool.HDRPSNR(hm1, hm2)
		ssim = hdrtool.HDRSSIM(hm1, hm2)
	} else {
		mse, snr, psnr, peak = hdrtool.PSNR(m1, m2)
		ssim = hdrtool.SSIM(m1, m2)
	}

	fmt.Printf("MSE: %.8f\n", mse)
	fmt.Printf("SNR: %.2f dB\n", snr)
	fmt.Printf("PSNR(max=%.4f): %.2f dB\n", peak, psnr)
	fmt.Printf("SSIM: %.8f\n", ssim)
	return nil
}
