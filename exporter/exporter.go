package exporter

import (
	"image"
	"image/jpeg"
	"os"
)

// Exports img to a local file found at path.
func ExportToLocalFile(img image.Image, path string) error {
	file, err := os.Create(path)
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

// Exports img to a cloud file in the given bucket at the given path.
func ExportToCloudFile(img image.Image, bucket string, path string) error {
	panic("Not implemented")
}
