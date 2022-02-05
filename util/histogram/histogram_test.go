package histogram

import (
	"testing"

	"fxkt.tech/mango/image"

	"fxkt.tech/mango"
	"fxkt.tech/mango/image/draw"
)

func TestLuma(t *testing.T) {
	infile := "../../test/images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	hiscanvas, err := Luma(canvas)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "../../test/images/out_effect.png"
	err = mango.WriteFile(hiscanvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDrawBlack(t *testing.T) {
	var canvas draw.Image = image.NewRGBA(image.Rect(0, 0, 256, 256))
	outfile := "../../test/images/black.jpg"
	err := mango.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}
