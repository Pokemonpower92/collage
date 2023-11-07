package calculator

import (
	"image"
	"image/color"
	"log"
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
