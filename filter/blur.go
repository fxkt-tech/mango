package filter

import (
	"image/color"
	"image/draw"
	"sync"

	"fxkt.tech/egami"
)

// warning: this implement is very slow, need to optimize.
// one goroutine: 5.44s
// goroutine for per-pix: 1.92s
// goroutine for per-line: 1.28s
func SlowBoxBlur(cvs draw.Image, radius, power int) error {
	if cvs == nil {
		return egami.ErrCanvasIsNil
	}

	rc := cvs.Bounds()

	w := rc.Max.X - rc.Min.X + 1
	h := rc.Max.Y - rc.Min.Y + 1

	for pw := 0; pw < power; pw++ {
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
		wg := &sync.WaitGroup{}
		for i := 0; i < w; i++ {
			wg.Add(1)
			go func(i int) {
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
				wg.Done()
			}(i)
		}
		wg.Wait()

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

	return nil
}

func BoxBlur(cvs draw.Image, radius, power int) error {
	if cvs == nil {
		return egami.ErrCanvasIsNil
	}

	// rc := cvs.Bounds()
	// w := rc.Max.X - rc.Min.X + 1
	// h := rc.Max.Y - rc.Min.Y + 1

	return nil
}
