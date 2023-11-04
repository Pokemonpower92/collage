package loader

import (
	"log"
	"testing"
)

func TestLoadLocalImage(t *testing.T) {
	path := "./testImages/gopher.png"
	if _, err := LoadLocalImage(path); err != nil {
		log.Printf("TestLoadLocalImage FAILED: %v\n", err)
		t.Fail()
	} else {
		t.Failed()
	}
}

func TestLoadLocalImageInvalidPath(t *testing.T) {
	path := "NoSuchPath"
	_, err := LoadLocalImage(path)
	if err == nil {
		log.Printf("TestLoadLocalImageInvalidPath FAILED: %v\n", err)
		t.Fail()
	}
}
