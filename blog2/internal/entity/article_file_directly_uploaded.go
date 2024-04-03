package entity

type ArticleFileUploadedID string

type ArticleFileUploaded struct {
	ID        ArticleFileUploadedID
	ArticleID ArticleID
	Type      ArticleFileType
}
