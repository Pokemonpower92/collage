package main

import (
	"fmt"
	"log/slog"

	"github.com/pokemonpower92/collage/colormap"
	"github.com/pokemonpower92/collage/common"
	"github.com/pokemonpower92/collage/creator"
	"github.com/pokemonpower92/collage/exporter"
	"github.com/pokemonpower92/collage/loader"
	"github.com/pokemonpower92/collage/settings"
)

func main() {
	imageLoaderSettings := settings.NewSettings()
	il := loader.NewImageLoader(*imageLoaderSettings)

	target, err := il.LoadTargetImage("images/test_images/target_images/gopher.png")
	if err != nil {
		slog.Error(fmt.Sprintf("%v\n", err))
	}

	if imageSet, err := il.LoadImageSet("images/image_sets/penis"); err != nil {
		slog.Error(fmt.Sprintf("%v\n", err))
	} else {
		slog.Info(fmt.Sprintf("Successfully loaded an image set.\n"))

		colormap := colormap.NewColorMap(imageSet)
		slog.Info(fmt.Sprintf("Successfully generated colormapping.\n"))

		collage := creator.Create(target, colormap, 6)
		slog.Info(fmt.Sprintf("Successfully generated collage.\n"))

		collage = common.Resize(collage, common.Dimensions{Height: 800, Width: 600})
		slog.Info(fmt.Sprintf("Successfully resized collage.\n"))

		if err := exporter.ExportToLocalFile(collage, "./output.jpeg"); err != nil {
			slog.Error(fmt.Sprintf("%v\n", err))
		}
	}
}
