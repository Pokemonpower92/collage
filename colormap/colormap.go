package colormap

import (
	"image"
	"image/color"
	"math"
)

// ColorMap maps the average colors of an image to
// that image.
type ColorMap struct {
	AverageColors []color.RGBA
	ImageSet      []*image.RGBA
}

// Calculate the 4th-dimensional distance between two colors.
func calculateColorDistance(one color.RGBA, two color.RGBA) float32 {
	r := math.Pow(float64(one.R)-float64(two.R), 2)
	g := math.Pow(float64(one.G)-float64(two.G), 2)
	b := math.Pow(float64(one.B)-float64(two.B), 2)
	a := math.Pow(float64(one.A)-float64(two.A), 2)

	combo := r + g + b + a

	return float32(math.Sqrt(combo))
}

// Finds the image from the image set with the
// closest average color to the target color.
func (colormap ColorMap) FindClosestImage(target color.RGBA) *image.RGBA {
	smallestDist := float32(math.MaxFloat32)
	closest := colormap.ImageSet[0]

	for idx, color := range colormap.AverageColors {
		dist := calculateColorDistance(color, target)
		if dist < smallestDist {
			smallestDist = dist
			closest = colormap.ImageSet[idx]
		}

	}

	return closest
}

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

// Creates a new ColorMap from the given image set.
func NewColorMap(imageSet []*image.RGBA) ColorMap {
	newMapping := ColorMap{}

	for _, img := range imageSet {
		newMapping.AverageColors = append(newMapping.AverageColors, calculateAverageColor(*img))
	}
	newMapping.ImageSet = imageSet
	return newMapping
}
