package encoder

import (
	"image"
	"image/gif"
	"io"

	"github.com/suzuito/sandbox2-go/common/terrors"
)

type EncoderGIF struct{}

func (t *EncoderGIF) Encode(w io.Writer, img image.Image) error {
	if err := gif.Encode(w, img, &gif.Options{}); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
func (t *EncoderGIF) GetMediaType() string {
	return "image/gif"
}
