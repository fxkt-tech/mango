package mango

import (
	"os"
	"path/filepath"

	"fxkt.tech/mango/errors"
	"fxkt.tech/mango/image"

	"fxkt.tech/mango/image/color"
	"fxkt.tech/mango/image/draw"
	"fxkt.tech/mango/image/jpeg"
	"fxkt.tech/mango/image/png"
)

// ReadFile is read image by file.
func ReadFile(path string) (draw.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	canvas := image.NewNRGBA(img.Bounds())

	rect := img.Bounds()
	for x := rect.Min.X; x <= rect.Max.X; x++ {
		for y := rect.Min.Y; y <= rect.Max.Y; y++ {
			p := img.At(x, y)
			r, g, b, a := p.RGBA()
			canvas.Set(x, y, &color.NRGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			})
		}
	}
	return canvas, nil
}

func WriteFile(cvs draw.Image, path string) (err error) {
	if cvs == nil {
		return errors.CanvasIsNil
	}

	file, err := os.Create(path)
	if err != nil {
		return
	}

	suffix := filepath.Ext(path)
	switch suffix {
	case ".jpeg", ".jpg":
		return jpeg.Encode(file, cvs, nil)
	case ".png":
		return png.Encode(file, cvs)
	}
	return errors.ExtNotSupported
}

func Copy(i draw.Image) draw.Image {
	rc := i.Bounds()
	o := image.NewRGBA(image.Rect(0, 0, rc.Dx(), rc.Dy()))
	for y := rc.Min.Y; y < rc.Max.Y; y++ {
		for x := rc.Min.X; x < rc.Max.X; x++ {
			r, g, b, a := i.At(x, y).RGBA()
			o.Set(x, y, color.NRGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			})
		}
	}
	return o
}
