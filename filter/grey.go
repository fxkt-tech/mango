package filter

import (
	"fxkt.tech/mango/errors"
	"fxkt.tech/mango/image/color"
	"fxkt.tech/mango/image/draw"
	"fxkt.tech/mango/math"
)

type GreyType uint8

var (
	GreyAvg     GreyType = 0 // 均度灰值
	GreyClassic GreyType = 1 // 经典灰值
	GreyPS      GreyType = 2 // PS灰值
)

func Grey(cvs draw.Image, gt GreyType) error {
	if cvs == nil {
		return errors.CanvasIsNil
	}

	rc := cvs.Bounds()
	for y := rc.Min.Y; y < rc.Max.Y; y++ {
		for x := rc.Min.X; x < rc.Max.X; x++ {
			r, g, b, a := cvs.At(x, y).RGBA()
			var grey uint32
			switch gt {
			case GreyAvg:
				grey = (r + g + b) / 3
			case GreyClassic:
				grey = uint32(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
			case GreyPS:
				max, min := math.MaxMin(r, g, b)
				grey = (max + min) / 2
			}
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

func BlackWhite(cvs draw.Image, thres uint8) error {
	if cvs == nil {
		return errors.CanvasIsNil
	}

	rc := cvs.Bounds()
	for y := rc.Min.Y; y < rc.Max.Y; y++ {
		for x := rc.Min.X; x < rc.Max.X; x++ {
			r, g, b, a := cvs.At(x, y).RGBA()
			var bw uint8
			if uint8(((r+g+b)>>8)/3) > thres {
				bw = 255
			} else {
				bw = 0
			}
			cvs.Set(x, y, &color.NRGBA{
				R: bw,
				G: bw,
				B: bw,
				A: uint8(a >> 8),
			})
		}
	}

	return nil
}

func Negative(cvs draw.Image) error {
	if cvs == nil {
		return errors.CanvasIsNil
	}

	rc := cvs.Bounds()
	for y := rc.Min.Y; y < rc.Max.Y; y++ {
		for x := rc.Min.X; x < rc.Max.X; x++ {
			r, g, b, a := cvs.At(x, y).RGBA()
			grey := uint32(0.3*float64(r) + 0.59*float64(g) + 0.11*float64(b))
			cvs.Set(x, y, &color.NRGBA{
				R: uint8(255 - grey>>8),
				G: uint8(255 - grey>>8),
				B: uint8(255 - grey>>8),
				A: uint8(a >> 8),
			})
		}
	}

	return nil
}

// 怀旧滤镜
// R = (0.393 * r + 0.769 * g + 0.189 * b)
// G = (0.349 * r + 0.686 * g + 0.168 * b)
// B = (0.272 * r + 0.534 * g + 0.131 * b)
// TODO: 会出现颜色不对的地方，可以再加一层grey滤镜（但是还是处理不好）
func ReAncient(cvs draw.Image) error {
	if cvs == nil {
		return errors.CanvasIsNil
	}

	rc := cvs.Bounds()
	for y := rc.Min.Y; y < rc.Max.Y; y++ {
		for x := rc.Min.X; x < rc.Max.X; x++ {
			r, g, b, a := cvs.At(x, y).RGBA()
			rr := uint32(0.393*float64(r) + 0.769*float64(g) + 0.189*float64(b))
			gg := uint32(0.349*float64(r) + 0.686*float64(g) + 0.168*float64(b))
			bb := uint32(0.272*float64(r) + 0.534*float64(g) + 0.131*float64(b))
			cvs.Set(x, y, &color.NRGBA{
				R: uint8(rr >> 8),
				G: uint8(gg >> 8),
				B: uint8(bb >> 8),
				A: uint8(a >> 8),
			})
		}
	}

	return nil
}

// 连环画滤镜（暂不可用）
// R = |g – b + g + r| * r / 256
// G = |b – g + b + r| * r / 256
// B = |b – g + b + r| * g / 256
func Comic(cvs draw.Image) error {
	if cvs == nil {
		return errors.CanvasIsNil
	}

	rc := cvs.Bounds()
	for y := rc.Min.Y; y < rc.Max.Y; y++ {
		for x := rc.Min.X; x < rc.Max.X; x++ {
			r, g, b, a := cvs.At(x, y).RGBA()
			rr := uint32(float64(math.Abs(int(g)-int(b)+int(g)+int(r))) * float64(r) / 256)
			gg := uint32(float64(math.Abs(int(b)-int(g)+int(b)+int(r))) * float64(r) / 256)
			bb := uint32(float64(math.Abs(int(b)-int(g)+int(b)+int(r))) * float64(g) / 256)
			cvs.Set(x, y, &color.NRGBA{
				R: uint8(rr >> 8),
				G: uint8(gg >> 8),
				B: uint8(bb >> 8),
				A: uint8(a >> 8),
			})
		}
	}

	return nil
}

// 抖音滤镜

// 浮雕滤镜 Engrave
