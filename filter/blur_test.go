package filter

import (
	"testing"

	"fxkt.tech/mango"
)

func TestMeanBlur(t *testing.T) {
	infile := "../test/images/ganyu4.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	err = MeanBlur(canvas, 2)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "../test/images/out_effect2.jpg"
	err = mango.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkMainBlur(b *testing.B) {
	infile := "../test/images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		MeanBlur(canvas, 10)
	}
}

func TestGaussBlur(t *testing.T) {
	infile := "../test/images/ganyu4.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	newcvs, err := GaussBlur(canvas, 5)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "../test/images/out_effect.jpg"
	err = mango.WriteFile(newcvs, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkGaussBlur(b *testing.B) {
	infile := "../test/images/ganyu4.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		GaussBlur(canvas, 5)
	}
	// 3605250100 ns/op	735910000 B/op	181445021 allocs/op
	// 3571080200 ns/op	735910128 B/op	181445020 allocs/op
	// 3751142700 ns/op	735910456 B/op	181445023 allocs/op
}

func TestSlowBoxBlur(t *testing.T) {
	infile := "../test/images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	err = SlowBoxBlur(canvas, 10, 2)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "../test/images/out_effect.jpg"
	err = mango.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkSlowBoxBlur(b *testing.B) {
	infile := "../test/images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		SlowBoxBlur(canvas, 10, 5)
	}
}
