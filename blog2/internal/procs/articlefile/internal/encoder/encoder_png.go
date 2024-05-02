package encoder

import (
	"image"
	"image/png"
	"io"

	"github.com/suzuito/sandbox2-go/common/terrors"
)

type EncoderPNG struct{}

func (t *EncoderPNG) Encode(w io.Writer, img image.Image) error {
	if err := png.Encode(w, img); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
func (t *EncoderPNG) GetMediaType() string {
	return "image/png"
}
