package usecase

import "github.com/suzuito/sandbox2-go/blog2/internal/usecase/internal/usecaseimpl"

type Usecase interface {
}

func NewUsecase() Usecase {
	return &usecaseimpl.UsecaseImpl{}
}
