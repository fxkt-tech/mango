package filter

import (
	"image/draw"

	"fxkt.tech/egami"
)

type RotateType uint8

var (
	Rotate180 RotateType = 0
)

func Rotate(cvs draw.Image, rt RotateType) error {
	if cvs == nil {
		return egami.ErrCanvasIsNil
	}

	// TODO: ...

	return nil
}
