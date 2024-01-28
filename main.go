package main

import (
	"fmt"
	"log/slog"

	"github.com/pokemonpower92/collage/colormap"
	"github.com/pokemonpower92/collage/creator"
	"github.com/pokemonpower92/collage/exporter"
	"github.com/pokemonpower92/collage/loader"
	"github.com/pokemonpower92/collage/settings"
)

func main() {
	imageLoaderSettings := settings.NewSettings()
	imageLoaderEnvironment := settings.NewEnvironment()

	il := loader.NewImageLoader(*imageLoaderSettings)

	target, err := il.LoadTargetImage("images/test_images/target_images/gopher.png")
	if err != nil {
		slog.Error(fmt.Sprintf("%v\n", err))
	}

	if imageSet, err := il.LoadImageSet("images/test_images/image_sets/rgba"); err != nil {
		slog.Error(fmt.Sprintf("%v\n", err))
	} else {
		slog.Info(fmt.Sprintf("Successfully loaded an image set.\n"))

		colormap := colormap.NewColorMap(imageSet)
		slog.Info(fmt.Sprintf("Successfully generated colormapping.\n"))

		ctr := creator.NewCreator(*imageLoaderSettings, *imageLoaderEnvironment)
		ctr.SetTargetImage(target)
		ctr.SetColorMap(&colormap)

		collage := ctr.Create()
		slog.Info(fmt.Sprintf("Successfully generated collage.\n"))

		if err := exporter.ExportToLocalFile(collage, "./output.jpeg"); err != nil {
			slog.Error(fmt.Sprintf("%v\n", err))
		}
	}
}
