package filter

import (
	"image/color"
	"image/draw"

	"fxkt.tech/egami"
)

func Ghoul(cvs draw.Image) error {
	if cvs == nil {
		return egami.ErrCanvasIsNil
	}

	rc := cvs.Bounds()
	for x := rc.Min.X; x <= rc.Max.X; x++ {
		for y := rc.Min.Y; y <= rc.Max.Y; y++ {
			r, g, b, a := cvs.At(x, y).RGBA()
			grey := uint32(0.3*float64(r) + 0.59*float64(g) + 0.11*float64(b))
			cvs.Set(x, y, &color.NRGBA{
				R: uint8(grey >> 8),
				G: uint8(grey >> 8),
				B: uint8(grey >> 8),
				A: uint8(a >> 8),
			})
		}
	}

	return nil
}
