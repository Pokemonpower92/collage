package loader

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/pokemonpower92/collage/internal/common"
	"github.com/pokemonpower92/collage/internal/settings"
)

func TestConvertToRGBA(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	rgba := convertToRGBA(img)
	if rgba.Bounds() != img.Bounds() {
		t.Errorf("Unexpected dimensions. Expected: %v, Got: %v", img.Bounds(), rgba.Bounds())
	}
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if rgba.At(x, y) != img.At(x, y) {
				t.Errorf("Unexpected pixel value at (%d, %d). Expected: %v, Got: %v", x, y, img.At(x, y), rgba.At(x, y))
			}
		}
	}
}

func TestLoadImage(t *testing.T) {
	testCases := []struct {
		name           string
		path           string
		dims           common.Dimensions
		expectedWidth  int
		expectedHeight int
	}{
		{
			name:           "Test case 1",
			path:           "../../images/test_images/target_images/gopher.png",
			dims:           common.Dimensions{Width: 200, Height: 200},
			expectedWidth:  200,
			expectedHeight: 200,
		},
		{
			name:           "Test case 2",
			path:           "../../images/test_images/target_images/gopher.png",
			dims:           common.Dimensions{Width: 300, Height: 300},
			expectedWidth:  300,
			expectedHeight: 300,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			img, err := LoadImage(tc.path, tc.dims)
			if err != nil {
				t.Fatalf("Failed to load image: %v", err)
			}
			if img.Bounds().Dx() != tc.expectedWidth || img.Bounds().Dy() != tc.expectedHeight {
				t.Errorf("Unexpected dimensions. Expected: %dx%d, Got: %dx%d",
					tc.expectedWidth,
					tc.expectedHeight,
					img.Bounds().Dx(),
					img.Bounds().Dy(),
				)
			}
		})
	}
}
func TestLoadImageSet(t *testing.T) {
	tempDir := t.TempDir()
	imagePaths := []string{
		filepath.Join(tempDir, "image1.png"),
		filepath.Join(tempDir, "image2.png"),
		filepath.Join(tempDir, "image3.png"),
	}
	for _, imagePath := range imagePaths {
		createTestImage(imagePath)
	}
	loader := &ImageLoader{
		Settings: *settings.NewSettings(),
	}
	imgSet, err := loader.LoadImageSet(tempDir)
	if err != nil {
		t.Fatalf("Failed to load image set: %v", err)
	}
	expectedNumImages := len(imagePaths)
	if len(imgSet) != expectedNumImages {
		t.Errorf("Unexpected number of loaded images. Expected: %d, Got: %d", expectedNumImages, len(imgSet))
	}
	for _, img := range imgSet {
		if img.Bounds().Dx() != loader.Settings.ImageSetDims.Width || img.Bounds().Dy() != loader.Settings.ImageSetDims.Height {
			t.Errorf("Unexpected dimensions for loaded image. Expected: %dx%d, Got: %dx%d",
				loader.Settings.ImageSetDims.Width,
				loader.Settings.ImageSetDims.Height,
				img.Bounds().Dx(),
				img.Bounds().Dy(),
			)
		}
	}
}

func createTestImage(path string) {
	img := image.NewRGBA(image.Rect(0, 0, 200, 200))
	saveImage(img, path)
}

func saveImage(img image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}
