package imageconverter

import (
	"image"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/procs/articlefile/pkg/encoder"
)

type ImageConverter interface {
	Decode(r io.Reader) (img image.Image, imgEncoder encoder.Encoder, thumbnailEncoder encoder.Encoder, err error)
	CreateThumbnail(src image.Image) image.Image
}
