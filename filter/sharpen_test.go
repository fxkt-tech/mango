package filter

import (
	"fmt"
	"testing"

	"fxkt.tech/mango"
)

func TestLaplaceSharpen(t *testing.T) {
	infile := "../test/images/ganyu5.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	newcanvas, err := LaplaceSharpen(canvas, 4)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "../test/images/out_effect.jpg"
	err = mango.WriteFile(newcanvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkLaplaceSharpen(b *testing.B) {
	infile := "../test/images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		LaplaceSharpen(canvas, 4)
	}
}

func TestXX(t *testing.T) {
	var a, b uint8 = 1, 2
	fmt.Println(a - b)
}
