package encode

import (
	"image/draw"
	"image/jpeg"
	"os"

	"fxkt.tech/egami"
)

func WriteFile(cvs draw.Image, path string) (err error) {
	if cvs == nil {
		return egami.ErrCanvasIsNil
	}

	file, err := os.Create(path)
	if err != nil {
		return
	}

	return jpeg.Encode(file, cvs, nil)
}
