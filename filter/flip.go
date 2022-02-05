package filter

import (
	"fxkt.tech/mango/errors"
	"fxkt.tech/mango/image/draw"
)

type FlipType uint8

const (
	HorizontalFlip FlipType = iota
	VerticalFlip
)

// Flip is ...
func Flip(cvs draw.Image, ft FlipType) error {
	if cvs == nil {
		return errors.CanvasIsNil
	}

	r := cvs.Bounds()

	switch ft {
	case HorizontalFlip:
		w := r.Max.X - r.Min.X
		for y := r.Min.Y; y <= r.Max.Y; y++ {
			for x := r.Min.X; x <= (r.Min.X+r.Max.X)/2; x++ {
				p1 := cvs.At(x, y)
				p2 := cvs.At(w-x, y)
				cvs.Set(x, y, p2)
				cvs.Set(w-x, y, p1)
			}
		}
	case VerticalFlip:
		h := r.Max.Y - r.Min.Y
		for y := r.Min.Y; y <= (r.Min.Y+r.Max.Y)/2; y++ {
			for x := r.Min.X; x <= r.Max.X; x++ {
				p1 := cvs.At(x, y)
				p2 := cvs.At(x, h-y)
				cvs.Set(x, y, p2)
				cvs.Set(x, h-y, p1)
			}
		}
	}

	return nil
}
