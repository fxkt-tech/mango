package test

import (
	"fmt"
	"os"
	"testing"

	"fxkt.tech/mango/image/jpeg"
)

func TestCodec(t *testing.T) {
	f, err := os.Open("src.jpeg")
	if err != nil {
		t.Error(err)
	}
	img, err := jpeg.Decode(f)
	if err != nil {
		t.Error(err)
	}
	f2, err := os.Create("dest.jpeg")
	if err != nil {
		t.Error(err)
	}
	err = jpeg.Encode(f2, img, nil)
}

func TestXX(t *testing.T) {
	x := 1 << 16
	fmt.Println(x)
}
