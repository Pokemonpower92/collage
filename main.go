package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log/slog"
	"os"

	"github.com/pokemonpower92/collage/colormap"
	"github.com/pokemonpower92/collage/creator"
	"github.com/pokemonpower92/collage/loader"
)

func saveImageToFile(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the image to the desired format (e.g., JPEG)
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	target, err := loader.LoadLocalImage("images/test_images/target_images/nancy.jpeg")
	if err != nil {
		slog.Error(fmt.Sprintf("%v\n", err))
	}

	if imageSet, err := loader.LoadLocalImageSet("images/test_images/image_sets/penis"); err != nil {
		slog.Error(fmt.Sprintf("%v\n", err))
	} else {
		slog.Info(fmt.Sprintf("Successfully loaded an image set with %d image(s).\n", len(imageSet)))

		colormap := colormap.NewColorMap(imageSet)
		slog.Info(fmt.Sprintf("ColorMap: %v\n", colormap))

		collage := creator.Create(target, colormap)

		if err := saveImageToFile(collage, "./output.jpeg"); err != nil {
			slog.Error(fmt.Sprintf("%v\n", err))
		}
	}
}
