package filter

import (
	"image"
	"image/draw"

	"fxkt.tech/egami"
)

// Scale is linear.
func Scale(cvs draw.Image, ow, oh int) (draw.Image, error) {
	if cvs == nil {
		return nil, egami.ErrCanvasIsNil
	}

	r := cvs.Bounds()
	iw := r.Max.X - r.Min.X
	ih := r.Max.Y - r.Min.Y
	newcvs := image.NewNRGBA(image.Rect(0, 0, ow, oh))
	for ox := 0; ox <= ow; ox++ {
		for oy := 0; oy <= oh; oy++ {
			ix := ox * iw / ow
			iy := oy * ih / oh
			c := cvs.At(ix, iy)
			newcvs.Set(ox, oy, c)
		}
	}

	return newcvs, nil
}
