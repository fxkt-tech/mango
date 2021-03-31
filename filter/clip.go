package filter

import (
	"image"
	"image/color"
	"image/draw"

	"fxkt.tech/egami"
)

// Clip is cut a area.
func Clip(cvs draw.Image, s image.Rectangle) (draw.Image, error) {
	if cvs == nil {
		return nil, egami.ErrCanvasIsNil
	}

	newcvs := image.NewNRGBA(image.Rect(0, 0, s.Max.X-s.Min.X, s.Max.Y-s.Min.Y))
	r := cvs.Bounds()
	rs := r.Intersect(s)
	for x := rs.Min.X; x <= rs.Max.X; x++ {
		for y := rs.Min.Y; y <= rs.Max.Y; y++ {
			r, g, b, a := cvs.At(x, y).RGBA()
			newcvs.Set(x-rs.Min.X, y-rs.Min.Y, &color.NRGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			})
		}
	}

	return newcvs, nil
}
