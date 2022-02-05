package test

import (
	"testing"

	"fxkt.tech/mango/image"

	"fxkt.tech/mango"
	"fxkt.tech/mango/filter"
)

func TestClip(t *testing.T) {
	infile := "images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	scalecanvas, err := filter.Clip(canvas, image.Rect(300, 300, 400, 400))
	if err != nil {
		t.Fatal(err)
	}

	outfile := "all_out.jpg"
	err = mango.WriteFile(scalecanvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFlip(t *testing.T) {
	infile := "images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	err = filter.Flip(canvas, filter.VerticalFlip)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "all_out.jpg"
	err = mango.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReAncient(t *testing.T) {
	infile := "images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	err = filter.ReAncient(canvas)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "all_out.jpg"
	err = mango.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestComic(t *testing.T) {
	infile := "images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	err = filter.Comic(canvas)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "all_out.jpg"
	err = mango.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestScale(t *testing.T) {
	infile := "images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	scalecanvas, err := filter.Scale(canvas, 40, 64)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "all_out.jpg"
	err = mango.WriteFile(scalecanvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSlowBoxBlur(t *testing.T) {
	infile := "images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	err = filter.SlowBoxBlur(canvas, 10, 1)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "../test/images/out_effect.jpg"
	err = mango.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGhoul(t *testing.T) {
	infile := "images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	err = filter.Ghoul(canvas)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "redleaf_ghoul.jpg"
	err = mango.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}
