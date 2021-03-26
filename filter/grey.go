package filter

import (
	"image/color"
	"image/draw"

	"fxkt.tech/egami"
)

type GreyType uint8

var (
	GreyAvg GreyType = 0
)

func Grey(cvs draw.Image, gt GreyType) error {
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

func BlackWhite(cvs draw.Image, thres uint8) error {
	if cvs == nil {
		return egami.ErrCanvasIsNil
	}

	rc := cvs.Bounds()
	for x := rc.Min.X; x <= rc.Max.X; x++ {
		for y := rc.Min.Y; y <= rc.Max.Y; y++ {
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
		return egami.ErrCanvasIsNil
	}

	rc := cvs.Bounds()
	for x := rc.Min.X; x <= rc.Max.X; x++ {
		for y := rc.Min.Y; y <= rc.Max.Y; y++ {
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
// TODO: 会出现颜色不对的地方，可以再加一层grey滤镜
func ReAncient(cvs draw.Image) error {
	if cvs == nil {
		return egami.ErrCanvasIsNil
	}

	rc := cvs.Bounds()
	for x := rc.Min.X; x <= rc.Max.X; x++ {
		for y := rc.Min.Y; y <= rc.Max.Y; y++ {
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

// 连环画滤镜
// R = |g – b + g + r| * r / 256
// G = |b – g + b + r| * r / 256
// B = |b – g + b + r| * g / 256
func Comic(cvs draw.Image) error {
	if cvs == nil {
		return egami.ErrCanvasIsNil
	}

	// rc := cvs.Bounds()
	// for x := rc.Min.X; x <= rc.Max.X; x++ {
	// 	for y := rc.Min.Y; y <= rc.Max.Y; y++ {
	// 		r, g, b, a := cvs.At(x, y).RGBA()
	// 		rr := uint32(0.393*float64(r) + 0.769*float64(g) + 0.189*float64(b))
	// 		gg := uint32(0.349*float64(r) + 0.686*float64(g) + 0.168*float64(b))
	// 		bb := uint32(0.272*float64(r) + 0.534*float64(g) + 0.131*float64(b))
	// 		cvs.Set(x, y, &color.NRGBA{
	// 			R: uint8(rr >> 8),
	// 			G: uint8(gg >> 8),
	// 			B: uint8(bb >> 8),
	// 			A: uint8(a >> 8),
	// 		})
	// 	}
	// }

	return nil
}

// 抖音滤镜

// 浮雕滤镜 Engrave
