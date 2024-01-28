package loader

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func TestConvertToRGBA(t *testing.T) {
	// Create a sample image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	// Call the ConvertToRGBA function
	rgba := convertToRGBA(img)

	// Check if the dimensions of the converted image match the original image
	if rgba.Bounds() != img.Bounds() {
		t.Errorf("Unexpected dimensions. Expected: %v, Got: %v", img.Bounds(), rgba.Bounds())
	}

	// Check if the pixels of the converted image match the original image
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if rgba.At(x, y) != img.At(x, y) {
				t.Errorf("Unexpected pixel value at (%d, %d). Expected: %v, Got: %v", x, y, img.At(x, y), rgba.At(x, y))
			}
		}
	}

	// Add more assertions or checks if needed
}

func TestResize(t *testing.T) {
	// Create a sample image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	// Define the expected dimensions of the resized image
	expectedWidth := 200
	expectedHeight := 200

	// Call the Resize function
	resized := Resize(img, Dimensions{Height: expectedHeight, Width: expectedWidth})

	// Check if the dimensions of the resized image match the expected dimensions
	if resized.Bounds().Dx() != expectedWidth || resized.Bounds().Dy() != expectedHeight {
		t.Errorf("Unexpected dimensions. Expected: %dx%d, Got: %dx%d", expectedWidth, expectedHeight, resized.Bounds().Dx(), resized.Bounds().Dy())
	}

	// Add more assertions or checks if needed
}

func TestLoadImage(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name           string
		path           string
		dims           Dimensions
		expectedWidth  int
		expectedHeight int
	}{
		{
			name:           "Test case 1",
			path:           "../images/test_images/target_images/gopher.png",
			dims:           Dimensions{Width: 200, Height: 200},
			expectedWidth:  200,
			expectedHeight: 200,
		},
		{
			name:           "Test case 2",
			path:           "../images/test_images/target_images/gopher.png",
			dims:           Dimensions{Width: 300, Height: 300},
			expectedWidth:  300,
			expectedHeight: 300,
		},
		// Add more test cases if needed
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the LoadImage function
			img, err := LoadImage(tc.path, tc.dims)
			if err != nil {
				t.Fatalf("Failed to load image: %v", err)
			}

			// Check if the dimensions of the loaded image match the expected dimensions
			if img.Bounds().Dx() != tc.expectedWidth || img.Bounds().Dy() != tc.expectedHeight {
				t.Errorf("Unexpected dimensions. Expected: %dx%d, Got: %dx%d", tc.expectedWidth, tc.expectedHeight, img.Bounds().Dx(), img.Bounds().Dy())
			}

			// Add more assertions or checks if needed
		})
	}
}
func TestLoadImageSet(t *testing.T) {
	// Create a temporary directory for test images
	tempDir := t.TempDir()

	// Create test images in the temporary directory
	imagePaths := []string{
		filepath.Join(tempDir, "image1.png"),
		filepath.Join(tempDir, "image2.png"),
		filepath.Join(tempDir, "image3.png"),
	}
	for _, imagePath := range imagePaths {
		createTestImage(imagePath)
	}

	// Create an instance of ImageLoader
	loader := &ImageLoader{
		ImageSetDims: Dimensions{Width: 200, Height: 200},
	}

	// Call the LoadImageSet function
	imgSet, err := loader.LoadImageSet(tempDir)
	if err != nil {
		t.Fatalf("Failed to load image set: %v", err)
	}

	// Check if the number of loaded images matches the expected number
	expectedNumImages := len(imagePaths)
	if len(imgSet) != expectedNumImages {
		t.Errorf("Unexpected number of loaded images. Expected: %d, Got: %d", expectedNumImages, len(imgSet))
	}

	// Check if the dimensions of each loaded image match the expected dimensions
	for _, img := range imgSet {
		if img.Bounds().Dx() != loader.ImageSetDims.Width || img.Bounds().Dy() != loader.ImageSetDims.Height {
			t.Errorf("Unexpected dimensions for loaded image. Expected: %dx%d, Got: %dx%d", loader.ImageSetDims.Width, loader.ImageSetDims.Height, img.Bounds().Dx(), img.Bounds().Dy())
		}
	}

	// Add more assertions or checks if needed
}

// Helper function to create a test image
func createTestImage(path string) {
	img := image.NewRGBA(image.Rect(0, 0, 200, 200))
	// Add code to modify the test image if needed
	saveImage(img, path)
}

// Helper function to save an image to disk
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
