package test

import (
	"testing"

	"fxkt.tech/egami/decode"
	"fxkt.tech/egami/encode"
	"fxkt.tech/egami/filter"
)

func TestAll(t *testing.T) {
	infile := "images/emma.jpg"
	canvas, err := decode.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	// scalecanvas, err := filter.Clip(canvas, image.Rect(300, 300, 400, 400))
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// err = filter.Flip(canvas, filter.VerticalFlip)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// err = filter.ReAncient(canvas)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// scalecanvas, err = filter.Scale(canvas, 40, 64)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// outfile := "all_out.jpg"
	// err = encode.WriteFile(scalecanvas, outfile)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	err = filter.SlowBoxBlur(canvas, 10, 5)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "all_out.jpg"
	err = encode.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGhoul(t *testing.T) {
	infile := "images/emma.jpg"
	canvas, err := decode.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	err = filter.Ghoul(canvas)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "redleaf_ghoul.jpg"
	err = encode.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}
