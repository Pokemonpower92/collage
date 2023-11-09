package main

import (
	"fmt"
	"log/slog"

	"github.com/pokemonpower92/collage/colormap"
	"github.com/pokemonpower92/collage/creator"
	"github.com/pokemonpower92/collage/exporter"
	"github.com/pokemonpower92/collage/loader"
)

func main() {
	target, err := loader.LoadLocalImage("images/test_images/target_images/gopher.png")
	if err != nil {
		slog.Error(fmt.Sprintf("%v\n", err))
	}

	if imageSet, err := loader.LoadLocalImageSet("images/test_images/image_sets/rgba"); err != nil {
		slog.Error(fmt.Sprintf("%v\n", err))
	} else {
		slog.Info(fmt.Sprintf("Successfully loaded an image set.\n"))

		colormap := colormap.NewColorMap(imageSet)
		slog.Info(fmt.Sprintf("Successfully generated colormapping.\n"))

		collage := creator.Create(target, colormap)

		if err := exporter.ExportToLocalFile(collage, "./output.jpeg"); err != nil {
			slog.Error(fmt.Sprintf("%v\n", err))
		}
	}
}
