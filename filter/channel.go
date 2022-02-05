package filter

import (
	"fxkt.tech/mango/errors"
	"fxkt.tech/mango/image/color"
	"fxkt.tech/mango/image/draw"
	"fxkt.tech/mango/math"
)

type ChannelType uint8

const (
	ChannelR ChannelType = iota
	ChannelG
	ChannelB
	ChannelA
	ChannelY
	ChannelCb
	ChannelCr
)

// SelectChannelsByRGB is select some color channel at rgb.
func SelectChannelsByRGB(cvs draw.Image, chls ...ChannelType) error {
	if cvs == nil {
		return errors.CanvasIsNil
	}

	r := cvs.Bounds()
	for y := r.Min.Y; y <= r.Max.Y; y++ {
		for x := r.Min.X; x <= r.Max.X; x++ {
			r, g, b, a := cvs.At(x, y).RGBA()
			var r32, g32, b32, a32 uint32
			for _, chl := range chls {
				switch chl {
				case ChannelR:
					r32 = r
				case ChannelG:
					g32 = g
				case ChannelB:
					b32 = b
				case ChannelA:
					a32 = a
				default:
					return errors.ChannelNotExist
				}
			}
			c := &color.NRGBA{
				R: uint8(r32 >> 8),
				G: uint8(g32 >> 8),
				B: uint8(b32 >> 8),
				A: uint8(a32 >> 8),
			}
			cvs.Set(x, y, c)
		}
	}

	return nil
}

// SelectChannelsByYUV is select some color channel at yuv.
func SelectChannelsByYUV(cvs draw.Image, chls ...ChannelType) error {
	if cvs == nil {
		return errors.CanvasIsNil
	}

	r := cvs.Bounds()
	for y := r.Min.Y; y <= r.Max.Y; y++ {
		for x := r.Min.X; x <= r.Max.X; x++ {
			clr := cvs.At(x, y).(color.NRGBA)
			y1, u, v := color.RGBToYCbCr(clr.R, clr.G, clr.B)
			var y8, u8, v8, a8 uint8
			for _, chl := range chls {
				switch chl {
				case ChannelY:
					y8 = y1
				case ChannelCb:
					u8 = u
				case ChannelCr:
					v8 = v
				case ChannelA:
					a8 = clr.A
				default:
					return errors.ChannelNotExist
				}
			}
			r8, g8, b8 := color.YCbCrToRGB(y8, u8, v8)
			c := &color.NRGBA{
				R: r8,
				G: g8,
				B: b8,
				A: a8,
			}
			cvs.Set(x, y, c)
		}
	}

	return nil
}

// 亮度/对比度调整算法
// avg: 图片灰度均值
// 当contrast>0时
// RGB = RGB + bright
// RGB = avg + (RGB-avg)/1-(contrast-100)
// 否则
// RGB = avg + (RGB-avg)*(1+contrast/100)
// RGB = RGB + bright
func BrightContrast(cvs draw.Image, bright, contrast int) error {
	bright = math.Clip(bright, -100, 100)
	contrast = math.Clip(contrast, -100, 100)
	var avg int
	rc := cvs.Bounds()
	for y := rc.Min.Y; y <= rc.Max.Y; y++ {
		for x := rc.Min.X; x <= rc.Max.X; x++ {
			clr := cvs.At(x, y).(color.NRGBA)
			avg = avg + (299*int(clr.R)+587*int(clr.G)+114*int(clr.B))/1000
		}
	}
	iw := rc.Max.X - rc.Min.X
	ih := rc.Max.Y - rc.Min.Y
	avg = avg / (iw * ih)

	bcmap := make([]int, 256)
	for i := 0; i < 256; i++ {
		var temp int
		if contrast > 0 {
			temp = math.Clip(i+bright, 0, 255)
			temp = math.Clip(avg+(temp-avg)*(1.0/(1.0-contrast/100.0)), 0, 255)
		} else {
			temp = i
			temp = math.Clip(avg+(temp-avg)*(1.0+contrast/100.0), 0, 255)
			temp = math.Clip(temp+bright, 0, 255)
		}
		bcmap[i] = temp
	}

	for y := rc.Min.Y; y <= rc.Max.Y; y++ {
		for x := rc.Min.X; x <= rc.Max.X; x++ {
			clr := cvs.At(x, y).(color.NRGBA)
			c := &color.NRGBA{
				R: uint8(bcmap[clr.R]),
				G: uint8(bcmap[clr.G]),
				B: uint8(bcmap[clr.B]),
				A: clr.A,
			}
			cvs.Set(x, y, c)
		}
	}
	return nil
}

// 饱和度算法（未调通）
func Saturation(cvs draw.Image, saturation int) error {
	if cvs == nil {
		return errors.CanvasIsNil
	}

	rc := cvs.Bounds()
	for y := rc.Min.Y; y <= rc.Max.Y; y++ {
		for x := rc.Min.X; x <= rc.Max.X; x++ {
			c := cvs.At(x, y).(color.NRGBA)
			max, min := math.MaxMin(c.R, c.G, c.B)
			delta, value := int(max-min), int(max)+int(min)
			if delta == 0 {
				continue
			}
			hslL := value >> 1
			var hslS int
			if delta != 0 {
				hslS = (delta << 7) / (5 - math.Abs(value-255))
			}
			k := saturation * 128 / 100
			var alpha int
			if k >= 0 {
				if k+hslS >= 128 {
					alpha = hslS
				} else {
					alpha = 128 - k
				}
				alpha = 128*128/alpha - 128
			} else {
				alpha = k
			}
			cvs.Set(x, y, &color.NRGBA{
				R: uint8(math.Clip(int(c.R)+((int(c.R)-hslL)*alpha>>7), 0, 255)),
				G: uint8(math.Clip(int(c.G)+((int(c.G)-hslL)*alpha>>7), 0, 255)),
				B: uint8(math.Clip(int(c.B)+((int(c.B)-hslL)*alpha>>7), 0, 255)),
				A: uint8(c.A),
			})
		}
	}

	return nil
}
