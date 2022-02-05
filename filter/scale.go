package filter

import (
	"fxkt.tech/mango/errors"
	"fxkt.tech/mango/image"

	"fxkt.tech/mango/image/draw"
)

// Scale is linear.
func Scale(cvs draw.Image, ow, oh int) (draw.Image, error) {
	if cvs == nil {
		return nil, errors.CanvasIsNil
	}

	r := cvs.Bounds()
	iw := r.Max.X - r.Min.X
	ih := r.Max.Y - r.Min.Y
	newcvs := image.NewNRGBA(image.Rect(0, 0, ow, oh))
	for oy := 0; oy <= oh; oy++ {
		for ox := 0; ox <= ow; ox++ {
			ix := ox * iw / ow
			iy := oy * ih / oh
			c := cvs.At(ix, iy)
			newcvs.Set(ox, oy, c)
		}
	}

	return newcvs, nil
}
