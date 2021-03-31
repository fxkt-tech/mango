package filter

import (
	"image/color"
	"image/draw"

	"fxkt.tech/egami"
)

// Ghoul is a lab.
func Ghoul(cvs draw.Image) error {
	if cvs == nil {
		return egami.ErrCanvasIsNil
	}

	radius := 10
	num := 5
	for i := 0; i < num; i++ {
		boxblur_one(cvs, radius)
	}

	return nil
}

func boxblur_one(cvs draw.Image, radius int) {
	rc := cvs.Bounds()

	w := rc.Max.X - rc.Min.X + 1
	h := rc.Max.Y - rc.Min.Y + 1

	var yy2d [][]uint8 // 原始Y
	for i := 0; i < w; i++ {
		yy2dh := make([]uint8, h)
		yy2d = append(yy2d, yy2dh)
	}
	for x := rc.Min.X; x <= rc.Max.X; x++ {
		for y := rc.Min.Y; y <= rc.Max.Y; y++ {
			r, g, b, _ := cvs.At(x, y).RGBA()
			yy, _, _ := color.RGBToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8))
			yy2d[x][y] = yy
		}
	}

	var dstyy2d [][]uint8 // box计算后的Y
	for i := 0; i < w; i++ {
		dstyy2dh := make([]uint8, h)
		dstyy2d = append(dstyy2d, dstyy2dh)
	}
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			var dstnum uint32
			var avalinum uint32
			for ir := i - radius; ir <= i+radius; ir++ {
				for jr := j - radius; jr <= j+radius; jr++ {
					if ir < 0 || ir >= w || jr < 0 || jr >= h {
						continue
					}
					avalinum = avalinum + 1
					dstnum = dstnum + uint32(yy2d[ir][jr])
				}
			}
			dstyy2d[i][j] = uint8(dstnum / avalinum)
		}
	}

	for x := rc.Min.X; x <= rc.Max.X; x++ {
		for y := rc.Min.Y; y <= rc.Max.Y; y++ {
			r, g, b, a := cvs.At(x, y).RGBA()
			_, cb, cr := color.RGBToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8))
			yy := dstyy2d[x][y]
			rr, gg, bb := color.YCbCrToRGB(yy, cb, cr)
			cvs.Set(x, y, &color.NRGBA{
				R: rr,
				G: gg,
				B: bb,
				A: uint8(a >> 8),
			})
		}
	}
}
