package usecase

import "context"

type Usecase interface {
	NotifyOnGCF(
		ctx context.Context,
		fullPath string,
	) error
}
