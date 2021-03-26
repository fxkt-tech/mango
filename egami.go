package egami

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

var ErrCanvasIsNil = errors.New("canvas is nil.")

type Image struct {
	canvas draw.Image
}

func NewImage() *Image {
	return &Image{}
}

// ReadFile is read image by file.
func (ig *Image) ReadFile(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return
	}

	ig.canvas = image.NewNRGBA(img.Bounds())

	rect := img.Bounds()
	for x := rect.Min.X; x <= rect.Max.X; x++ {
		for y := rect.Min.Y; y <= rect.Max.Y; y++ {
			p := img.At(x, y)
			r, g, b, a := p.RGBA()
			ig.canvas.Set(x, y, &color.NRGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			})
		}
	}
	return
}

type ChannelType uint8

var (
	RedChannel   ChannelType = 0
	GreenChannel ChannelType = 1
	BlueChannel  ChannelType = 2

	ErrChannelNotExist = errors.New("channel not exist.")
)

// FilterChannel is drop some color channel.
func (ig *Image) FilterChannel(cts ...ChannelType) error {
	if ig.canvas == nil {
		return ErrCanvasIsNil
	}

	return ig.yljz(func(x, y int, cvs draw.Image) error {
		p := cvs.At(x, y)
		r, g, b, a := p.RGBA()
		c := &color.NRGBA{
			R: uint8(r >> 8),
			G: uint8(g >> 8),
			B: uint8(b >> 8),
			A: uint8(a >> 8),
		}
		for _, ct := range cts {
			switch ct {
			case RedChannel:
				c.R = uint8(0)
			case GreenChannel:
				c.G = uint8(0)
			case BlueChannel:
				c.B = uint8(0)
			default:
				return ErrChannelNotExist
			}
		}
		cvs.Set(x, y, c)
		return nil
	})
}

// Clip is cut a area.
func (ig *Image) Clip(s image.Rectangle) error {
	if ig.canvas == nil {
		return ErrCanvasIsNil
	}

	// TODO: ...
	canvas := image.NewNRGBA(image.Rect(0, 0, s.Max.X-s.Min.X, s.Max.Y-s.Min.Y))
	r := ig.canvas.Bounds()
	rs := r.Intersect(s)
	fmt.Println(r, s, rs)
	for x := rs.Min.X; x <= rs.Max.X; x++ {
		for y := rs.Min.Y; y <= rs.Max.Y; y++ {
			p := ig.canvas.At(x, y)
			r, g, b, a := p.RGBA()
			canvas.Set(x-rs.Min.X, y-rs.Min.Y, &color.NRGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			})
		}
	}
	ig.canvas = canvas

	return nil
}

type FlipType uint8

var (
	HorizontalFlip FlipType = 0
	VerticalFlip   FlipType = 1
)

// Flip is
func (ig *Image) Flip(ft FlipType) error {
	if ig.canvas == nil {
		return ErrCanvasIsNil
	}

	rect := ig.canvas.Bounds()

	switch ft {
	case HorizontalFlip:
		width := rect.Max.X - rect.Min.X
		for x := rect.Min.X; x <= (rect.Min.X+rect.Max.X)/2; x++ {
			for y := rect.Min.Y; y <= rect.Max.Y; y++ {
				p1 := ig.canvas.At(x, y)
				p2 := ig.canvas.At(width-x, y)
				ig.canvas.Set(x, y, p2)
				ig.canvas.Set(width-x, y, p1)
			}
		}
	case VerticalFlip:
	}

	return nil
}

type RotateType uint8

var (
	Rotate180 RotateType = 0
)

func (ig *Image) Rotate(rt RotateType) error {
	if ig.canvas == nil {
		return ErrCanvasIsNil
	}

	// TODO: ...

	return nil
}

func (ig *Image) Write(path string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}

	return jpeg.Encode(file, ig.canvas, nil)
}

type Traverse func(x, y int, cvs draw.Image) error

func (ig *Image) yljz(f Traverse) (err error) {
	if ig.canvas == nil {
		return ErrCanvasIsNil
	}
	rect := ig.canvas.Bounds()
	for x := rect.Min.X; x <= rect.Max.X; x++ {
		for y := rect.Min.Y; y <= rect.Max.Y; y++ {
			err = f(x, y, ig.canvas)
			if err != nil {
				return
			}
		}
	}
	return nil
}
