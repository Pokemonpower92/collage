package settings

import (
	"github.com/pokemonpower92/collage/common"
)

type Settings struct {
	ImageSetDims    common.Dimensions
	TargetImageDims common.Dimensions
	FinalImageDims  common.Dimensions
}

func NewSettings() *Settings {
	return &Settings{
		ImageSetDims:    common.Dimensions{Height: 100, Width: 100},
		TargetImageDims: common.Dimensions{Height: 100, Width: 100},
		FinalImageDims:  common.Dimensions{Height: 100, Width: 100},
	}
}
