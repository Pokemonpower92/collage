package loader

import (
	"image/color"
	"log"
	"testing"
)

const targetPath = "../images/test_images/target_images/gopher.png"
const imageSetPath = "../images/test_images/target_images"
const badPath = "Bad Path"

func TestLoadLocalImage(t *testing.T) {
	if img, err := LoadLocalImage(targetPath); err != nil {
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
	_, err := LoadLocalImage(badPath)
	if err == nil {
		log.Printf("TestLoadLocalImageInvalidPath FAILED: %v\n", err)
		t.Fail()
	}
}

func TestLoadLocalImageSet(t *testing.T) {
	if imgs, err := LoadLocalImageSet(imageSetPath); err != nil {
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
	if _, err := LoadLocalImageSet(badPath); err == nil {
		log.Printf("TestLoadLocalImageSetInvalidPath FAILED: %v\n", err)
		t.Fail()
	}

	t.Failed()
}
