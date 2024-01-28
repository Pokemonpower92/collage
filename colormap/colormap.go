package colormap

import (
	"image"
	"image/color"
	"math"
)

// calculateAverageColor calculates the average color of an image.
// It takes an image.RGBA as input and returns the average color as a color.RGBA struct.
// The average color is calculated by iterating through all pixels of the image,
// summing up the color values, and dividing them by the total number of pixels.
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

// calculateColorDistance calculates the Euclidean distance between two RGBA colors.
// It takes two color.RGBA values as input and returns the distance as a float32 value.
func calculateColorDistance(one color.RGBA, two color.RGBA) float32 {
	// Calculate the squared differences for each color channel
	r := math.Pow(float64(one.R)-float64(two.R), 2)
	g := math.Pow(float64(one.G)-float64(two.G), 2)
	b := math.Pow(float64(one.B)-float64(two.B), 2)
	a := math.Pow(float64(one.A)-float64(two.A), 2)

	// Calculate the sum of squared differences
	combo := r + g + b + a

	// Return the square root of the sum as the distance
	return float32(math.Sqrt(combo))
}

// ColorMap represents a collection of average colors and image sets.
type ColorMap struct {
	AverageColors []color.RGBA
	ImageSet      []*image.RGBA
}

// NewColorMap creates a new ColorMap.
// It initializes the AverageColors and ImageSet fields with empty slices.
func NewColorMap() ColorMap {
	return ColorMap{
		AverageColors: make([]color.RGBA, 0),
		ImageSet:      make([]*image.RGBA, 0),
	}
}

// AddImage adds an image to the ColorMap.
// It appends the image to the ImageSet and calculates the average color of the image,
// which is then appended to the AverageColors slice.
func (colormap *ColorMap) AddImage(img *image.RGBA) {
	colormap.ImageSet = append(colormap.ImageSet, img)
	colormap.AverageColors = append(colormap.AverageColors, calculateAverageColor(*img))
}

// AddImages adds multiple images to the ColorMap.
// It takes a slice of *image.RGBA as input and calls AddImage for each image in the slice.
func (colormap *ColorMap) AddImages(imgs []*image.RGBA) {
	for _, img := range imgs {
		colormap.AddImage(img)
	}
}

// FindClosestImage finds the closest image in the colormap to the target color.
// It calculates the color distance between each average color in the colormap and the target color,
// and returns the image with the smallest color distance.
// If the colormap is empty, it returns nil.
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
