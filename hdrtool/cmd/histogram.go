package cmd

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"time"

	// Import LDR codecs
	_ "image/jpeg"
	"image/png"

	"github.com/mdouchement/hdr"
	"github.com/mdouchement/hdrtool"
	// Import HDR codecs
	_ "github.com/mdouchement/hdr/crad"
	_ "github.com/mdouchement/hdr/pfm"
	_ "github.com/mdouchement/hdr/rgbe"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	// HistogramCommand defines the command for compute luminance histogram.
	HistogramCommand = &cobra.Command{
		Use:   "histogram source_file",
		Short: "Performs luminance histogram and save it in the src folder",
		Long:  "Performs luminance histogram and save it in the src folder",
		RunE:  histogramAction,
	}
)

func histogramAction(c *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("histogram: Invalid number of arguments")
	}
	fi, err := os.Open(args[0])
	if err != nil {
		return errors.Wrap(err, "quality:")
	}
	defer fi.Close()

	start := time.Now()
	m, fname, err := image.Decode(fi)
	if err != nil {
		return errors.Wrap(err, "quality:")
	}
	fmt.Printf("Read image (%dx%dp - %s - %v) %s\n", m.Bounds().Dx(), m.Bounds().Dy(), fname, time.Since(start), filepath.Base(args[0]))

	hdrm, ok := m.(hdr.Image)
	if !ok {
		hdrm = hdrtool.NewLDRWrapper(m)
	}

	hist := hdrtool.Histogram(filepath.Base(args[0]), hdrm)

	fo, err := os.Create(fmt.Sprintf("%s.hist.png", args[0]))
	if err != nil {
		return errors.Wrap(err, "histogram:")
	}
	defer fo.Close()

	if err = png.Encode(fo, hist); err != nil {
		return errors.Wrap(err, "histogram:")
	}
	return fo.Sync()
}
