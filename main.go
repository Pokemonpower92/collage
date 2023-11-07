package main

import (
	"fmt"
	"log/slog"

	"github.com/pokemonpower92/collage/calculator"
	"github.com/pokemonpower92/collage/loader"
)

func main() {

	if imageSet, err := loader.LoadLocalImageSet("loader/test_images/image_sets/gopher"); err != nil {
		slog.Error(fmt.Sprintf("%v\n", err))
	} else {
		slog.Info(fmt.Sprintf("Successfully loaded an image set with %d image(s).\n", len(imageSet)))

		averages := calculator.CalculateImageSetAverages(imageSet)
		slog.Info(fmt.Sprintf("Found image set averages: %v\n", averages))
	}
}
