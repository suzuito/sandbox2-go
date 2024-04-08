package articlefile

import (
	internal_imageconverter "github.com/suzuito/sandbox2-go/blog2/internal/procs/articlefile/internal/imageconverter"
	"github.com/suzuito/sandbox2-go/blog2/internal/procs/articlefile/pkg/imageconverter"
)

func NewImageConverter() imageconverter.ImageConverter {
	return &internal_imageconverter.ImageConverter{}
}
