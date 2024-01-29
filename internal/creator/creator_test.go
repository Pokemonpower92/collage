package creator

import (
	"image"
	"testing"

	"github.com/pokemonpower92/collage/internal/colormap"
	"github.com/pokemonpower92/collage/internal/common"
	"github.com/pokemonpower92/collage/internal/settings"
)

func NewTestEnvironment() settings.Environment {
	return settings.Environment{
		CreatorThreads: 1,
	}
}

func TestCreateCollage(t *testing.T) {
	targetImage := image.NewRGBA(image.Rect(0, 0, 100, 100))

	colorMap := colormap.NewColorMap()
	colorMap.AddImage(image.NewRGBA(image.Rect(0, 0, 10, 10)))
	colorMap.AddImage(image.NewRGBA(image.Rect(0, 0, 10, 10)))
	colorMap.AddImage(image.NewRGBA(image.Rect(0, 0, 10, 10)))

	settings := settings.Settings{
		ImageSetDims: common.Dimensions{
			Width:  10,
			Height: 10,
		},
		FinalImageDims: common.Dimensions{
			Width:  100,
			Height: 100,
		},
	}

	creator := NewCreator(settings, NewTestEnvironment())
	creator.SetTargetImage(targetImage)
	creator.SetColorMap(&colorMap)

	collage := creator.Create()

	expectedCollageBounds := image.Rect(0, 0, 100, 100)
	if !collage.Bounds().Eq(expectedCollageBounds) {
		t.Errorf("Expected collage bounds to be %v, but got %v", expectedCollageBounds, collage.Bounds())
	}
}
