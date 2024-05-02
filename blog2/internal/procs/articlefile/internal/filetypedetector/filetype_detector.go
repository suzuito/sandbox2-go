package filetypedetector

import (
	"net/http"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type FileTypeDetector struct{}

func (t *FileTypeDetector) Do(data []byte) (entity.FileType, string) {
	mimeTypeDetected := http.DetectContentType(data)
	return entity.NewFileTypeFromMimeType(mimeTypeDetected), mimeTypeDetected
}
