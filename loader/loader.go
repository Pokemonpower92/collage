package loader

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"path/filepath"

	"golang.org/x/image/draw"
)

type Dimensions struct {
	Height int
	Width  int
}

type ImageLoadError struct {
	What string
}

func (ile *ImageLoadError) Error() string {
	return ile.What
}

func convertToRGBA(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	return rgba
}

// Resize the given img to the given dims
func Resize(img image.Image, dims Dimensions) *image.RGBA {
	bounds := image.Rect(0, 0, dims.Width, dims.Height)
	resized := image.NewRGBA(bounds)

	draw.ApproxBiLinear.Scale(resized, resized.Rect, img, img.Bounds(), draw.Over, nil)

	return resized
}

// Loads an image file indicated by the local path.
func LoadImage(path string, dims Dimensions) (*image.RGBA, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, &ImageLoadError{
			fmt.Sprintf("Failed to open image file: %s with error: %v", path, err),
		}
	}
	defer reader.Close()

	if img, _, err := image.Decode(reader); err != nil {
		return nil, &ImageLoadError{
			fmt.Sprintf("Failed to decode image file: %s with error: %v", path, err),
		}

	} else {
		return convertToRGBA(Resize(img, dims)), nil
	}
}

type ImageLoader struct {
	TargetDims   Dimensions
	ImageSetDims Dimensions
}

func NewImageLoader(targetDims Dimensions, imageSetDims Dimensions) *ImageLoader {
	return &ImageLoader{
		TargetDims:   targetDims,
		ImageSetDims: imageSetDims,
	}
}

// Loads the target image located in the given local path.
func (il *ImageLoader) LoadTargetImage(path string) (*image.RGBA, error) {
	img, err := LoadImage(path, il.TargetDims)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// Loads all images located in the given local path.
func (il *ImageLoader) LoadImageSet(path string) ([]*image.RGBA, error) {
	if imageFiles, err := os.ReadDir(path); err != nil {
		return nil, &ImageLoadError{
			fmt.Sprintf("Failed to open image set directory: %s with error: %v", path, err),
		}
	} else {
		imgSet := make([]*image.RGBA, 0)

		for _, imgFile := range imageFiles {
			imgFilePath := filepath.Join(path, imgFile.Name())
			if img, err := LoadImage(imgFilePath, il.ImageSetDims); err != nil {
				return nil, err
			} else {
				imgSet = append(imgSet, img)
			}
		}

		return imgSet, nil
	}

}
