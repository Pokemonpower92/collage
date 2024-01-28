package settings

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	ImageSetBucket     string
	TargtetImageBucket string
	CreatorThreads     int
}

// NewEnvironment creates a new instance of the Environment struct.
// It loads environment variables from the "../env" file using godotenv.Load.
// If there is an error loading the file, it panics with the message "Error loading .env file".
// It returns a pointer to the created Environment struct.
func NewEnvironment() *Environment {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	creatorThreads, err := strconv.Atoi(os.Getenv("CREATOR_THREADS"))
	if err != nil {
		panic("Invalid value for CREATOR_THREADS")
	}

	return &Environment{
		ImageSetBucket:     os.Getenv("IMAGE_SET_BUCKET"),
		TargtetImageBucket: os.Getenv("TARGET_IMAGE_BUCKET"),
		CreatorThreads:     creatorThreads,
	}
}
