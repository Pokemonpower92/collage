package loader

import (
	"image/color"
	"log"
	"testing"
)

func TestLoadLocalImage(t *testing.T) {
	path := "./test_images/target_images/gopher.png"
	if img, err := LoadLocalImage(path); err != nil {
		log.Printf("TestLoadLocalImage FAILED: %v\n", err)
		t.Fail()
	} else {
		if img == nil {
			t.Fail()
		}
	}
	t.Failed()
}

func TestLoadLocalImageInvalidPath(t *testing.T) {
	path := "NoSuchPath"
	_, err := LoadLocalImage(path)
	if err == nil {
		log.Printf("TestLoadLocalImageInvalidPath FAILED: %v\n", err)
		t.Fail()
	}
}

func TestLoadLocalImageSet(t *testing.T) {
	path := "./test_images/image_sets/gopher"
	if imgs, err := LoadLocalImageSet(path); err != nil {
		log.Printf("TestLoadLocalImageSet FAILED: %v\n", err)
		t.Fail()
	} else {
		if len(imgs) != 1 {
			t.Fail()
		}
		expectedColor := color.RGBA{0, 0, 0, 0}
		if foundColor := imgs[0].At(0, 0).(color.RGBA); foundColor != expectedColor {
			t.Fail()
		}
	}
	t.Failed()
}

func TestLoadLocalImageSetInvalidPath(t *testing.T) {
	invalidImageSetpath := "./test_images/image_set/gopher"
	if _, err := LoadLocalImageSet(invalidImageSetpath); err == nil {
		log.Printf("TestLoadLocalImageSetInvalidPath FAILED: %v\n", err)
		t.Fail()
	}

	t.Failed()
}
