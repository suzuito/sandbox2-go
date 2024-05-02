package articlefile

import (
	internal_filetypedetector "github.com/suzuito/sandbox2-go/blog2/internal/procs/articlefile/internal/filetypedetector"
	"github.com/suzuito/sandbox2-go/blog2/internal/procs/articlefile/pkg/filetypedetector"
)

func NewFileTypeDetector() filetypedetector.FileTypeDetector {
	return &internal_filetypedetector.FileTypeDetector{}
}
