package filetypedetector

import "github.com/suzuito/sandbox2-go/blog2/internal/entity"

type FileTypeDetector interface {
	Do(data []byte) (entity.FileType, string)
}
