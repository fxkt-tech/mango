package filter

import (
	"image"
	"image/draw"

	"fxkt.tech/egami"
	"fxkt.tech/egami/encode"
)

func Scale(cvs draw.Image, ow, oh int) error {
	if cvs == nil {
		return egami.ErrCanvasIsNil
	}

	// TODO: ...

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

	encode.WriteFile(newcvs, "xx.jpg")

	return nil
}
