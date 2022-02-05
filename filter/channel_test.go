package filter

import (
	"testing"

	"fxkt.tech/mango"
)

func TestSelectChannelsByRGB(t *testing.T) {
	infile := "../test/images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	err = SelectChannelsByRGB(canvas, ChannelB, ChannelA)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "../test/images/out_effect.jpg"
	err = mango.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSelectChannelsByYUV(t *testing.T) {
	infile := "../test/images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	err = SelectChannelsByYUV(canvas, ChannelY, ChannelCb, ChannelCr, ChannelA)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "../test/images/out_effect.jpg"
	err = mango.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBrightContrast(t *testing.T) {
	infile := "../test/images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	err = BrightContrast(canvas, -50, -50)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "../test/images/out_effect.jpg"
	err = mango.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSaturation(t *testing.T) {
	infile := "../test/images/xinhai.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		t.Fatal(err)
	}

	err = Saturation(canvas, 99)
	if err != nil {
		t.Fatal(err)
	}

	outfile := "../test/images/out_effect.jpg"
	err = mango.WriteFile(canvas, outfile)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkBrightContrast(b *testing.B) {
	infile := "../test/images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BrightContrast(canvas, 50, 50)
	}
}

func BenchmarkSaturation(b *testing.B) {
	infile := "../test/images/emma.jpg"
	canvas, err := mango.ReadFile(infile)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Saturation(canvas, 50)
	}
}
