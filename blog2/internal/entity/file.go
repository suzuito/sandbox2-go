package entity

import (
	"regexp"
	"time"
)

type FileID string

type FileType string

const (
	FileTypeImage   FileType = "image"
	FileTypeVideo   FileType = "video"
	FileTypeText    FileType = "text"
	FileTypeUnknown FileType = "unknown"
)

func NewFileTypeFromMimeType(mimeType string) FileType {
	// image
	reMimeTypeImage := regexp.MustCompile("image/*")
	if reMimeTypeImage.MatchString(mimeType) {
		return FileTypeImage
	}
	// video
	reMimeTypeVideo := regexp.MustCompile("video/*")
	if reMimeTypeVideo.MatchString(mimeType) {
		return FileTypeVideo
	}
	// text
	reMimeTypeText := regexp.MustCompile("text/*")
	if reMimeTypeText.MatchString(mimeType) {
		return FileTypeText
	}
	return FileTypeUnknown
}

type File struct {
	ID        FileID    `json:"id"`
	Type      FileType  `json:"type"`
	MediaType string    `json:"mediaType"` // Ex) application/json; charset=utf8
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type FileThumbnailID string

type FileThumbnail struct {
	ID        FileThumbnailID `json:"id"`
	FileID    FileID          `json:"fileId"`
	MediaType string          `json:"mediaType"` // Ex) application/json; charset=utf8
}

type FileAndThumbnail struct {
	File          *File          `json:"file"`
	FileThumbnail *FileThumbnail `json:"fileThumbnail"`
}
