package common

import (
	"image"
	"testing"
)

func TestResize(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	testCases := []struct {
		name           string
		img            image.Image
		dims           Dimensions
		expectedWidth  int
		expectedHeight int
	}{
		{
			name:           "Test case 1",
			img:            img,
			dims:           Dimensions{Width: 50, Height: 50},
			expectedWidth:  50,
			expectedHeight: 50,
		},
		{
			name:           "Test case 2",
			img:            img,
			dims:           Dimensions{Width: 200, Height: 200},
			expectedWidth:  200,
			expectedHeight: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resized := Resize(tc.img, tc.dims)

			if resized.Bounds().Dx() != tc.expectedWidth || resized.Bounds().Dy() != tc.expectedHeight {
				t.Errorf("Unexpected dimensions. Expected: %dx%d, Got: %dx%d", tc.expectedWidth, tc.expectedHeight, resized.Bounds().Dx(), resized.Bounds().Dy())
			}
		})
	}
}
