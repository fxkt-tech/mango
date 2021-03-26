package egami

import (
	"image"
	"testing"
	"time"
)

func TestFilterChannel(t *testing.T) {
	img := NewImage()
	err := img.ReadFile("redleaf.jpg")
	if err != nil {
		t.Fatal(err)
	}
	err = img.FilterChannel(RedChannel)
	if err != nil {
		t.Fatal(err)
	}
	err = img.Write("redleaf_filterchannel.jpg")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFlip(t *testing.T) {
	img := NewImage()
	err := img.ReadFile("redleaf.jpg")
	if err != nil {
		t.Fatal(err)
	}
	st := time.Now()
	err = img.Flip(HorizontalFlip)
	diff := time.Since(st)
	t.Log(diff)
	if err != nil {
		t.Fatal(err)
	}
	err = img.Write("redleaf_hflip.jpg")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClip(t *testing.T) {
	img := NewImage()
	err := img.ReadFile("redleaf.jpg")
	if err != nil {
		t.Fatal(err)
	}
	st := time.Now()
	err = img.Clip(image.Rect(0, 0, 4288, 2848))
	diff := time.Since(st)
	t.Log(diff)
	if err != nil {
		t.Fatal(err)
	}
	err = img.Write("redleaf_clip.jpg")
	if err != nil {
		t.Fatal(err)
	}
}
