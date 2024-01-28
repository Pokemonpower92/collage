package common

import (
	"image"

	"golang.org/x/image/draw"
)

// Resize the given img to the given dims
func Resize(img image.Image, dims Dimensions) *image.RGBA {
	bounds := image.Rect(0, 0, dims.Width, dims.Height)
	resized := image.NewRGBA(bounds)

	draw.ApproxBiLinear.Scale(resized, resized.Rect, img, img.Bounds(), draw.Over, nil)

	return resized
}
