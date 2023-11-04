package loader

import (
	"fmt"
	"image"
	"os"
	"time"
)

type ImageLoadError struct {
	When time.Time
	Err  error
}

func (ile *ImageLoadError) Error() string {
	return fmt.Sprintf("%v : Fatal error loading image: %s", ile.When, ile.Err)
}

func LoadCloudImage(bucket string, path string) ([]image.Image, error) {
	panic("Not Implemented")
}

// Loads an image file indicated by the local path.
func LoadLocalImage(path string) (image.Image, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, &ImageLoadError{
			time.Now(),
			err,
		}
	}
	defer reader.Close()

	if img, _, err := image.Decode(reader); err != nil {
		return img, nil
	} else {
		return nil, &ImageLoadError{
			time.Now(),
			err,
		}
	}
}

// Loads all images located in the given local path.
func LoadLocalImageSet(path string) ([]image.Image, error) {
	panic("Not implemented.")
}
