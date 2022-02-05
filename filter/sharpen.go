package filter

import (
	"fxkt.tech/mango/errors"
	"fxkt.tech/mango/image"
	"fxkt.tech/mango/image/color"
	"fxkt.tech/mango/image/draw"
	"fxkt.tech/mango/math"
)

// 拉普拉斯锐化，4邻域（锐化效果很不好，有横纵抽丝情况出现）
func LaplaceSharpen(cvs draw.Image, mode int) (draw.Image, error) {
	if cvs == nil {
		return nil, errors.CanvasIsNil
	}

	rc := cvs.Bounds()
	newcvs := image.NewNRGBA(image.Rect(0, 0, rc.Dx(), rc.Dy()))
	for y := rc.Min.Y; y < rc.Max.Y; y++ {
		for x := rc.Min.X; x < rc.Max.X; x++ {
			// 4邻域
			center := cvs.At(x, y).(color.NRGBA)
			top := cvs.At(x, math.Clip(y-1, 0, 255)).(color.NRGBA)
			bottom := cvs.At(x, math.Clip(y+1, 0, 255)).(color.NRGBA)
			left := cvs.At(math.Clip(x-1, 0, 255), y).(color.NRGBA)
			right := cvs.At(math.Clip(x+1, 0, 255), y).(color.NRGBA)
			r := math.Clip(int(center.R)+int(center.R)*4-int(top.R)-int(bottom.R)-int(left.R)-int(right.R), 0, 255)
			g := math.Clip(int(center.G)+int(center.G)*4-int(top.G)-int(bottom.G)-int(left.G)-int(right.G), 0, 255)
			b := math.Clip(int(center.B)+int(center.B)*4-int(top.B)-int(bottom.B)-int(left.B)-int(right.B), 0, 255)

			// 8邻域
			// center := cvs.At(x, y).(color.NRGBA)
			// top := cvs.At(x, math.Clip(y-1, 0, 255)).(color.NRGBA)
			// bottom := cvs.At(x, math.Clip(y+1, 0, 255)).(color.NRGBA)
			// left := cvs.At(math.Clip(x-1, 0, 255), y).(color.NRGBA)
			// right := cvs.At(math.Clip(x+1, 0, 255), y).(color.NRGBA)
			// topleft := cvs.At(math.Clip(x-1, 0, 255), math.Clip(y-1, 0, 255)).(color.NRGBA)
			// topright := cvs.At(math.Clip(x+1, 0, 255), math.Clip(y-1, 0, 255)).(color.NRGBA)
			// bottomleft := cvs.At(math.Clip(x-1, 0, 255), math.Clip(y+1, 0, 255)).(color.NRGBA)
			// bottomright := cvs.At(math.Clip(x+1, 0, 255), math.Clip(y+1, 0, 255)).(color.NRGBA)
			// r := math.Clip(int(center.R)+int(center.R)*8-int(top.R)-int(bottom.R)-int(left.R)-int(right.R)-int(topleft.R)-int(topright.R)-int(bottomleft.R)-int(bottomright.R), 0, 255)
			// g := math.Clip(int(center.G)+int(center.G)*8-int(top.G)-int(bottom.G)-int(left.G)-int(right.G)-int(topleft.G)-int(topright.G)-int(bottomleft.G)-int(bottomright.G), 0, 255)
			// b := math.Clip(int(center.B)+int(center.B)*8-int(top.B)-int(bottom.B)-int(left.B)-int(right.B)-int(topleft.B)-int(topright.B)-int(bottomleft.B)-int(bottomright.B), 0, 255)
			newcvs.Set(x, y, &color.NRGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				A: center.A,
			})
		}
	}

	return newcvs, nil
}
