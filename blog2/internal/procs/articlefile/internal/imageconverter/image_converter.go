package imageconverter

import (
	"image"
	"io"

	internal_encoder "github.com/suzuito/sandbox2-go/blog2/internal/procs/articlefile/internal/encoder"
	"github.com/suzuito/sandbox2-go/blog2/internal/procs/articlefile/pkg/encoder"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"golang.org/x/image/draw"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type ImageConverter struct {
}

func (t *ImageConverter) CreateThumbnail(src image.Image) image.Image {
	smallWidth := 50
	smallHeight := 50
	thumbnail := image.NewNRGBA(image.Rect(0, 0, smallWidth, smallHeight))
	draw.NearestNeighbor.Scale(
		thumbnail, thumbnail.Bounds(),
		src, src.Bounds(),
		draw.Over,
		nil,
	)
	return thumbnail
}

func (t *ImageConverter) Decode(r io.Reader) (image.Image, encoder.Encoder, encoder.Encoder, error) {
	src, imageFormat, err := image.Decode(r)
	if err != nil {
		return nil, nil, nil, terrors.Wrap(err)
	}
	var encoder encoder.Encoder
	encoderThumbnail := internal_encoder.EncoderJPEG{}
	switch imageFormat {
	case "jpeg":
		encoder = &internal_encoder.EncoderJPEG{}
	case "png":
		encoder = &internal_encoder.EncoderPNG{}
	case "gif":
		encoder = &internal_encoder.EncoderGIF{}
	default:
		return nil, nil, nil, terrors.Wrapf("unknown format %s", imageFormat)
	}
	return src, encoder, &encoderThumbnail, nil
}
