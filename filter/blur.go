package filter

import (
	sysmath "math"

	"fxkt.tech/mango/errors"
	"fxkt.tech/mango/image"
	"fxkt.tech/mango/image/color"
	"fxkt.tech/mango/image/draw"
	"fxkt.tech/mango/math"
)

// warning: this implement is very slow, need to optimize.
// one goroutine: 5.44s
// goroutine for per-pix: 1.92s
// goroutine for per-line: 1.28s
func SlowBoxBlur(cvs draw.Image, radius, power int) error {
	if cvs == nil {
		return errors.CanvasIsNil
	}

	rc := cvs.Bounds()
	w := rc.Dx()
	h := rc.Dy()
	for pw := 0; pw < power; pw++ {
		var yy2d [][]uint8 // 原始Y
		for i := 0; i < w; i++ {
			yy2dh := make([]uint8, h)
			yy2d = append(yy2d, yy2dh)
		}
		for x := rc.Min.X; x < rc.Max.X; x++ {
			for y := rc.Min.Y; y < rc.Max.Y; y++ {
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
		for j := 0; j < h; j++ {
			for i := 0; i < w; i++ {
				var dstnum uint32
				var avalinum uint32
				for jr := j - radius; jr <= j+radius; jr++ {
					for ir := i - radius; ir <= i+radius; ir++ {
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

		for y := rc.Min.Y; y < rc.Max.Y; y++ {
			for x := rc.Min.X; x < rc.Max.X; x++ {
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

	return nil
}

func MeanBlur(cvs draw.Image, radius int) error {
	if cvs == nil {
		return errors.CanvasIsNil
	}

	rc := cvs.Bounds()
	m := (2*radius + 1) * (2*radius + 1)
	for y := rc.Min.Y; y < rc.Max.Y; y++ {
		for x := rc.Min.X; x < rc.Max.X; x++ {
			// 以下为核函数
			var sumr, sumg, sumb int
			for n := -radius; n <= radius; n++ {
				for m := -radius; m <= radius; m++ {
					ny := math.Clip(y+n, 0, rc.Dy()-1)
					nx := math.Clip(x+m, 0, rc.Dx()-1)
					c := cvs.At(nx, ny).(color.NRGBA)
					sumr = sumr + int(c.R)
					sumg = sumg + int(c.G)
					sumb = sumb + int(c.B)
				}
			}
			cvs.Set(x, y, color.NRGBA{
				R: uint8(sumr / m),
				G: uint8(sumg / m),
				B: uint8(sumb / m),
				A: cvs.At(x, y).(color.NRGBA).A,
			})
		}
	}

	return nil
}

// 高斯模糊（效果没有ps好,不知道为啥。。。）
func GaussBlur(cvs draw.Image, radius int) (draw.Image, error) {
	if cvs == nil {
		return nil, errors.CanvasIsNil
	}

	rc := cvs.Bounds()
	newcvs := image.NewNRGBA(image.Rect(0, 0, rc.Dx(), rc.Dy()))
	stride := 2*radius + 1
	corefuncvals := gausscore(radius, float64(radius))
	for y := rc.Min.Y; y < rc.Max.Y; y++ {
		for x := rc.Min.X; x < rc.Max.X; x++ {
			// 以下为核函数
			var sumr, sumg, sumb float64
			for n := -radius; n <= radius; n++ {
				for m := -radius; m <= radius; m++ {
					ny := math.Clip(y+n, 0, rc.Dy()-1)
					nx := math.Clip(x+m, 0, rc.Dx()-1)
					c := cvs.At(nx, ny).(color.NRGBA)
					sumr = sumr + float64(c.R)*corefuncvals[(m+radius)+(n+radius)*stride]
					sumg = sumg + float64(c.G)*corefuncvals[(m+radius)+(n+radius)*stride]
					sumb = sumb + float64(c.B)*corefuncvals[(m+radius)+(n+radius)*stride]
				}
			}
			newcvs.Set(x, y, color.NRGBA{
				R: uint8(sumr),
				G: uint8(sumg),
				B: uint8(sumb),
				A: cvs.At(x, y).(color.NRGBA).A,
			})
		}
	}

	return newcvs, nil
}

func gausscore(radius int, sigma float64) (values []float64) {
	stride := 2*radius + 1
	var sum float64
	values = make([]float64, stride*stride)
	for y, h := -radius, 0; y <= radius; y, h = y+1, h+1 {
		for x, w := -radius, 0; x <= radius; x, w = x+1, w+1 {
			value := (1.0 / (2.0 * sysmath.Pi * sigma * sigma)) * sysmath.Exp(-(float64(x)*float64(x)+float64(y)*float64(y))/(2.0*sigma*sigma))
			values[w+h*stride] = value
			sum = sum + value
		}
	}

	for i := 0; i < stride*stride; i++ {
		values[i] = values[i] / sum
	}
	return
}
