package decode

import (
	"image"
	"image/color"
	"image/draw"
	"os"
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
