package colormap

import (
	"image"
	"image/color"
	"log"
	"math"
	"testing"
)

func newRedImage() *image.RGBA {
	bounds := image.Rect(0, 0, 100, 100)
	red := image.NewRGBA(bounds)

	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			red.Set(x, y, color.RGBA{255, 0, 0, 0})
		}
	}

	return red
}

func TestCalculateAverageColor(t *testing.T) {
	red := newRedImage()
	average := calculateAverageColor(*red)

	expectedColor := color.RGBA{255, 0, 0, 0}
	if average != expectedColor {
		log.Printf("TestCalculateAverageColor FAILED")
		t.Fail()
	}
	t.Failed()
}

func TestCaluculateColorDistance(t *testing.T) {
	one := color.RGBA{255, 0, 0, 0}
	two := color.RGBA{0, 255, 0, 0}

	expectedValue := float32(math.Sqrt(math.Pow(255, 2) * 2))

	if dist := calculateColorDistance(one, two); dist != expectedValue {
		log.Printf("TestCaluculateColorDistance FAILED. Dist: %f Expected: %f", dist, expectedValue)
		t.Fail()
	}

	t.Failed()
}
