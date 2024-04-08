package entity

import "mime"

type ArticleFileID string

type ArticleFileType string

const (
	ArticleFileTypeImage ArticleFileType = "image"
)

type ArticleFile struct {
	ID        ArticleFileID
	Type      ArticleFileType
	MediaType string // Ex) application/json; charset=utf8
}

func (t *ArticleFile) ExtIncludingDot() string {
	exts, err := mime.ExtensionsByType(t.MediaType)
	if err != nil {
		return ""
	}
	if len(exts) <= 0 {
		return ""
	}
	return exts[0]
}

type ArticleFileThumbnail struct {
	ID        ArticleFileID
	MediaType string // Ex) application/json; charset=utf8
}

func (t *ArticleFileThumbnail) ExtIncludingDot() string {
	exts, err := mime.ExtensionsByType(t.MediaType)
	if err != nil {
		return ""
	}
	if len(exts) <= 0 {
		return ""
	}
	return exts[0]
}
