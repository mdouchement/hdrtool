package hdrtool

import (
	"image"
	"sort"

	"github.com/mdouchement/hdr"
	chart "github.com/wcharczuk/go-chart"
)

const unit = 0.01

// Histogram create a Luminance histogram.
func Histogram(title string, m hdr.Image) image.Image {
	histogram := map[int]int{}
	for y := 0; y < m.Bounds().Dy(); y++ {
		for x := 0; x < m.Bounds().Dx(); x++ {
			_, Y, _, _ := m.HDRAt(x, y).HDRXYZA()

			histogram[scale(Y, unit)]++
		}
	}

	// Convert to graph format
	keys := make([]int, 0, len(histogram))
	for k := range histogram {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// Setup axes
	xv := make([]float64, len(histogram))
	yv := make([]float64, len(histogram))
	for i, k := range keys {
		xv[i] = unscale(k, unit)
		yv[i] = float64(histogram[k])
	}

	// Build histogram
	graph := chart.Chart{
		Title:      title,
		TitleStyle: chart.StyleShow(),
		XAxis: chart.XAxis{
			Name:      "Luminance",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "Level",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.ColorBlue,
					FillColor:   chart.ColorBlue.WithAlpha(100),
				},
				XValues: xv,
				YValues: yv,
			},
		},
	}

	collector := &chart.ImageWriter{}
	graph.Render(chart.PNG, collector)

	image, err := collector.Image()
	if err != nil {
		panic(err) // Should never be reached
	}
	return image
}
