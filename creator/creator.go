package creator

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log/slog"
	"sync"

	"github.com/pokemonpower92/collage/colormap"
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

// Create the collage of target from the imageSet
func Create(target *image.RGBA, colormap colormap.ColorMap, numThreads int) *image.RGBA {
	targetBounds := target.Bounds()
	collageBounds := image.Rect(0, 0, targetBounds.Dx()*350, targetBounds.Dy()*350)

	collage := image.NewRGBA(collageBounds)

	totalPixels := targetBounds.Dx() * targetBounds.Dy()
	pixelsPerThread := totalPixels / numThreads

	var wg sync.WaitGroup

	wg.Add(numThreads)

	for thread := 0; thread < numThreads; thread++ {
		start := thread * pixelsPerThread
		end := (thread + 1) * pixelsPerThread

		if thread == numThreads-1 {
			end = totalPixels
		}

		drawClosestImage := func(threadID int) {
			defer wg.Done()
			slog.Info(fmt.Sprintf("Entering thread  %d\n", threadID))
			for pixel := start; pixel != end; pixel++ {
				x := pixel % targetBounds.Dy()
				y := pixel / targetBounds.Dy()

				closestImage := colormap.FindClosestImage(target.At(x, y).(color.RGBA))
				destRect := image.Rect(x*350+350, y*350+350, x*350, y*350)
				// if threadID == 0 {
				// 	slog.Info(fmt.Sprintf("Drawing pixel %d\n", pixel))
				// }
				draw.Draw(collage, destRect, closestImage, image.Point{0, 0}, draw.Over)
				// if threadID == 0 {
				// 	slog.Info(fmt.Sprintf("Done drawing pixel %d\n", pixel))
				// }
			}
		}

		go drawClosestImage(thread)
	}

	wg.Wait()
	return collage
}
