package creator

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log/slog"
	"sync"

	"collage/colormap"
	"collage/common"
	"collage/settings"
)

// Creator is a struct that contains the settings and environment for creating a collage.
type Creator struct {
	settings    settings.Settings
	environment settings.Environment
	targetImage *image.RGBA
	colorMap    *colormap.ColorMap
}

// NewCreator creates a new instance of the Creator struct.
func NewCreator(settings settings.Settings, environment settings.Environment) *Creator {
	return &Creator{
		settings:    settings,
		environment: environment,
	}
}

// SetTargetImage sets the target image for the collage.
func (c *Creator) SetTargetImage(targetImage *image.RGBA) {
	c.targetImage = targetImage
}

// SetColorMap sets the color map for the collage.
func (c *Creator) SetColorMap(colorMap *colormap.ColorMap) {
	c.colorMap = colorMap
}

// Create generates a collage by finding the closest image for each pixel in the target image.
// It uses multiple threads to parallelize the process and improve performance.
// The collage is then resized to the specified dimensions before being returned.
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
				x := pixel % targetBounds.Dx()
				y := pixel / targetBounds.Dx()

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
