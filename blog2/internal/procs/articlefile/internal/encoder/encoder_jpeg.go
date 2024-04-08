package encoder

import (
	"image"
	"image/jpeg"
	"io"

	"github.com/suzuito/sandbox2-go/common/terrors"
)

type EncoderJPEG struct {
}

func (t *EncoderJPEG) Encode(w io.Writer, img image.Image) error {
	if err := jpeg.Encode(w, img, &jpeg.Options{
		Quality: 80,
	}); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}
func (t *EncoderJPEG) GetMediaType() string {
	return "image/jpeg"
}
