package settings

import (
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	imageSetBucket     string
	targtetImageBucket string
}

// NewEnvironment creates a new instance of the Environment struct.
// It loads environment variables from the "../env" file using godotenv.Load.
// If there is an error loading the file, it panics with the message "Error loading .env file".
// It returns a pointer to the created Environment struct.
func NewEnvironment() *Environment {
	err := godotenv.Load("../env")
	if err != nil {
		panic("Error loading .env file")
	}

	return &Environment{
		imageSetBucket:     os.Getenv("IMAGE_SET_BUCKET"),
		targtetImageBucket: os.Getenv("TARGET_IMAGE_BUCKET"),
	}
}
