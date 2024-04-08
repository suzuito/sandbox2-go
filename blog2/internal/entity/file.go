package entity

import "mime"

type FileID string

type FileType string

const (
	FileTypeImage FileType = "image"
)

type File struct {
	ID              FileID
	Name            string
	Type            FileType
	MediaType       string // Ex) application/json; charset=utf8
	ExistsThumbnail bool
}

func (t *File) ExtIncludingDot() string {
	exts, err := mime.ExtensionsByType(t.MediaType)
	if err != nil {
		return ""
	}
	if len(exts) <= 0 {
		return ""
	}
	return exts[0]
}

type FileThumbnail struct {
	ID        FileID
	MediaType string // Ex) application/json; charset=utf8
}

func (t *FileThumbnail) ExtIncludingDot() string {
	exts, err := mime.ExtensionsByType(t.MediaType)
	if err != nil {
		return ""
	}
	if len(exts) <= 0 {
		return ""
	}
	return exts[0]
}
