package filter

import (
	"fxkt.tech/mango/errors"
	"fxkt.tech/mango/image/draw"
)

type RotateType uint8

const (
	Rotate180 RotateType = iota
)

func Rotate(cvs draw.Image, rt RotateType) error {
	if cvs == nil {
		return errors.CanvasIsNil
	}

	// TODO: ...

	return nil
}
