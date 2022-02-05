package errors

import "errors"

var (
	CanvasIsNil     = errors.New("canvas is nil.")
	ChannelNotExist = errors.New("channel not exist.")
	ExtNotSupported = errors.New("ext is not supported.")
)
