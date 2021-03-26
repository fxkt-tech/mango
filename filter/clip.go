package filter

import (
	"image"
	"image/color"
	"image/draw"

	"fxkt.tech/egami"
)

// Clip is cut a area.
func Clip(cvs draw.Image, s image.Rectangle) error {
	if cvs == nil {
		return egami.ErrCanvasIsNil
	}

	canvas := image.NewNRGBA(image.Rect(0, 0, s.Max.X-s.Min.X, s.Max.Y-s.Min.Y))
	r := cvs.Bounds()
	rs := r.Intersect(s)
	for x := rs.Min.X; x <= rs.Max.X; x++ {
		for y := rs.Min.Y; y <= rs.Max.Y; y++ {
			p := cvs.At(x, y)
			r, g, b, a := p.RGBA()
			canvas.Set(x-rs.Min.X, y-rs.Min.Y, &color.NRGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			})
		}
	}
	cvs = canvas

	return nil
}
