package creator

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/pokemonpower92/collage/colormap"
)

// Create the collage of target from the imageSet
func Create(target *image.RGBA, colormap colormap.ColorMap) *image.RGBA {
	targetBounds := target.Bounds()
	collageBounds := image.Rect(0, 0, targetBounds.Dx()*200, targetBounds.Dy()*200)

	collage := image.NewRGBA(collageBounds)

	// For each pixel in the target image.
	for x := 0; x < targetBounds.Dx(); x++ {
		for y := 0; y < targetBounds.Dy(); y++ {
			closestImage := colormap.FindClosestImage(target.At(x, y).(color.RGBA))
			destRect := image.Rect(x*200+200, y*200+200, x*200, y*200)
			draw.Draw(collage, destRect, closestImage, image.Point{0, 0}, draw.Over)
		}
	}

	return collage
}
