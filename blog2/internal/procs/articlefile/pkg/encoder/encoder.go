package encoder

import (
	"image"
	"io"
)

type Encoder interface {
	Encode(w io.Writer, img image.Image) error
	GetMediaType() string
}
