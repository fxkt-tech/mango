package filter

import (
	"errors"
	"image/color"
	"image/draw"

	"fxkt.tech/egami"
)

type ChannelType uint8

var (
	ChannelRed   ChannelType = 0
	ChannelGreen ChannelType = 1
	ChannelBlue  ChannelType = 2

	ErrChannelNotExist = errors.New("channel not exist.")
)

// DropChannel is drop some color channel.
func DropChannel(cvs draw.Image, cts ...ChannelType) error {
	if cvs == nil {
		return egami.ErrCanvasIsNil
	}

	r := cvs.Bounds()
	for x := r.Min.X; x <= r.Max.X; x++ {
		for y := r.Min.Y; y <= r.Max.Y; y++ {
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
				case ChannelRed:
					c.R = uint8(0)
				case ChannelGreen:
					c.G = uint8(0)
				case ChannelBlue:
					c.B = uint8(0)
				default:
					return ErrChannelNotExist
				}
			}
			cvs.Set(x, y, c)
		}
	}

	return nil
}
