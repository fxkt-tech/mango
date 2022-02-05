package filter

import (
	"testing"

	"fxkt.tech/mango"
)

func TestGrey(t *testing.T) {
	infile := "../test/images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	err = Grey(canvas, GreyClassic)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "../test/images/out_effect.jpg"
	err = mango.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}
