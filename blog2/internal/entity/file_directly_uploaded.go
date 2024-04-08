package entity

type FileUploadedID string

type FileProcessStatus string

const (
	FileProcessStatusRegistered FileProcessStatus = "registered"
	FileProcessStatusError      FileProcessStatus = "error"
	FileProcessStatusDone       FileProcessStatus = "done"
)

type FileUploaded struct {
	ID            FileUploadedID
	Name          string
	Type          FileType
	ProcessStatus FileProcessStatus
}
