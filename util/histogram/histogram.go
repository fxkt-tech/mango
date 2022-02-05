package histogram

import (
	"fxkt.tech/mango/errors"
	"fxkt.tech/mango/image"
	"fxkt.tech/mango/math"

	"fxkt.tech/mango/image/color"
	"fxkt.tech/mango/image/draw"
)

// 亮度
func Luma(cvs draw.Image) (draw.Image, error) {
	if cvs == nil {
		return nil, errors.CanvasIsNil
	}

	// 计算图片所有red在0-255的数量
	yy := make([]int, 256)
	b := cvs.Bounds()
	x, y := b.Dx(), b.Dy()
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			r, g, b, _ := cvs.At(i, j).RGBA()
			r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)
			y8, _, _ := color.RGBToYCbCr(r8, g8, b8)
			yy[y8] = yy[y8] + 1
		}
	}
	// 创建256*256的canvas
	canvas := image.NewRGBA(image.Rect(0, 0, 255, 255))
	// 绘制坐标系
	for i, y := range yy {
		PaintFirstQuadrant(canvas, i, int(y)*255/math.MaxInSlice(yy...), color.RGBA{
			R: uint8(i), G: uint8(i), B: uint8(i), A: uint8(255),
		})
	}
	return canvas, nil
}

func PaintFirstQuadrant(cvs draw.Image, hx, hy int, c color.Color) {
	b := cvs.Bounds()
	for i := b.Dy() - hy - 1; i < b.Dy(); i++ {
		cvs.Set(hx, i, c)
	}
}
