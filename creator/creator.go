package creator

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log/slog"
	"sync"

	"github.com/pokemonpower92/collage/colormap"
	"github.com/pokemonpower92/collage/common"
	"github.com/pokemonpower92/collage/settings"
)

// Create: Will split the target image into equal parts
// and then start a new thread for each pointison.

// Each thread will run drawClosestImage which takes the dimensions of
// the split and will draw the closest images in place.

// --- --- --- --- -- -- -- --
// 20 / 8 = 2
// 20 % 8 = 4

// func drawClosestImage(collage *image.RGBA, target *image.RGBA, start int, end int, colormap colormap.ColorMap) {
// 	bounds := target.Bounds()

// 	for pixel := start; pixel != end; pixel++ {
// 		x := pixel % bounds.Dy()
// 		y := pixel / bounds.Dy()

// 		closestImage := colormap.FindClosestImage(target.At(x, y).(color.RGBA))
// 		destRect := image.Rect(x*350+350, y*350+350, x*350, y*350)
// 		draw.Draw(collage, destRect, closestImage, image.Point{0, 0}, draw.Over)
// 	}
// }

type Creator struct {
	settings    settings.Settings
	environment settings.Environment
	targetImage *image.RGBA
	colorMap    *colormap.ColorMap
}

func NewCreator(settings settings.Settings, environment settings.Environment) *Creator {
	return &Creator{
		settings:    settings,
		environment: environment,
	}
}

func (c *Creator) SetTargetImage(targetImage *image.RGBA) {
	c.targetImage = targetImage
}

func (c *Creator) SetColorMap(colorMap *colormap.ColorMap) {
	c.colorMap = colorMap
}

// Create the collage of target from the imageSet
func (c *Creator) Create() *image.RGBA {
	targetBounds := c.targetImage.Bounds()
	imageSetDims := c.settings.ImageSetDims

	collageBounds := image.Rect(0,
		0,
		targetBounds.Dx()*imageSetDims.Width,
		targetBounds.Dy()*imageSetDims.Height,
	)

	collage := image.NewRGBA(collageBounds)

	totalPixels := targetBounds.Dx() * targetBounds.Dy()
	pixelsPerThread := totalPixels / c.environment.CreatorThreads

	var wg sync.WaitGroup

	wg.Add(c.environment.CreatorThreads)

	slog.Info(fmt.Sprintf("Generating collage with %d threads.\n", c.environment.CreatorThreads))
	for thread := 0; thread < c.environment.CreatorThreads; thread++ {
		start := thread * pixelsPerThread
		end := (thread + 1) * pixelsPerThread

		if thread == c.environment.CreatorThreads-1 {
			end = totalPixels
		}

		drawClosestImage := func(threadID int) {
			defer wg.Done()
			for pixel := start; pixel != end; pixel++ {
				x := pixel % targetBounds.Dy()
				y := pixel / targetBounds.Dy()

				closestImage := c.colorMap.FindClosestImage(c.targetImage.At(x, y).(color.RGBA))
				destRect := image.Rect(x*imageSetDims.Width+imageSetDims.Width,
					y*imageSetDims.Height+imageSetDims.Height,
					x*imageSetDims.Width,
					y*imageSetDims.Height,
				)

				draw.Draw(collage, destRect, closestImage, image.Point{0, 0}, draw.Over)
			}
		}

		go drawClosestImage(thread)
	}

	wg.Wait()
	slog.Info(fmt.Sprintf("Successfully generated collage.\n"))

	slog.Info(fmt.Sprintf("Resizing collage to: %x.\n", c.settings.FinalImageDims))
	collage = common.Resize(collage, c.settings.FinalImageDims)
	slog.Info(fmt.Sprintf("Successfully resized collage.\n"))

	return collage
}
