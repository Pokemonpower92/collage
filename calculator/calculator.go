package calculator

import (
	"image"
	"image/color"
)

// Returns the average color of an image.
func calculateAverageColor(img image.RGBA) color.RGBA {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Initialize variables to accumulate color values
	var totalR, totalG, totalB, totalA uint32

	// Iterate through all pixels to calculate the sum of color values
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixelColor := img.At(x, y).(color.RGBA)
			totalR += uint32(pixelColor.R)
			totalG += uint32(pixelColor.G)
			totalB += uint32(pixelColor.B)
			totalA += uint32(pixelColor.A)
		}
	}

	// Calculate the average color values
	totalPixels := uint32(width * height)
	avgR := totalR / totalPixels
	avgG := totalG / totalPixels
	avgB := totalB / totalPixels
	avgA := totalA / totalPixels

	// Create the average color as a color.RGBA struct
	return color.RGBA{R: uint8(avgR), G: uint8(avgG), B: uint8(avgB), A: uint8(avgA)}
}

// Creates a slice of average colors for each image in
// an image set.
func CalculateImageSetAverages(imageSet []*image.RGBA) []color.RGBA {
	aveColors := make([]color.RGBA, 0)
	for _, img := range imageSet {
		aveColors = append(aveColors, calculateAverageColor(*img))
	}

	return aveColors
}
