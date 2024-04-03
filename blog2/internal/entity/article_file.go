package entity

type ArticleFileID string

type ArticleFileType string

const (
	ArticleFileTypeImage ArticleFileType = "image"
)

type ArticleFile struct {
	ID   ArticleFileID
	Type ArticleFileType
}
