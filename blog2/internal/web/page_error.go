package web

type PageError struct {
	ComponentCommonHead ComponentCommonHead
	Message             string
}

func NewPageErrorUnknownError() *PageError {
	return &PageError{
		ComponentCommonHead: ComponentCommonHead{},
		Message:             "謎のエラーが発生した！",
	}
}

func NewPageErrorNotFound() *PageError {
	return &PageError{
		ComponentCommonHead: ComponentCommonHead{},
		Message:             "not found",
	}
}
