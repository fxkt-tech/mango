package test

import (
	"testing"

	"fxkt.tech/egami/decode"
	"fxkt.tech/egami/encode"
	"fxkt.tech/egami/filter"
)

func TestAll(t *testing.T) {
	infile := "../redleaf.jpg"
	canvas, err := decode.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	// err = filter.Clip(canvas, image.Rect(0, 0, 100, 100))
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// err = filter.Flip(canvas, filter.VerticalFlip)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	err = filter.ReAncient(canvas)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "redleaf_out.jpg"
	err = encode.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGhoul(t *testing.T) {
	infile := "../redleaf.jpg"
	canvas, err := decode.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	err = filter.ReAncient(canvas)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "redleaf_ghoul.jpg"
	err = encode.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}
