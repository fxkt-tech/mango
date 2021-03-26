package util

import (
	"image/draw"

	"fxkt.tech/egami"
)

func MergeImage(cvs, layer draw.Image) error {
	if cvs == nil || layer == nil {
		return egami.ErrCanvasIsNil
	}
	return nil
}
