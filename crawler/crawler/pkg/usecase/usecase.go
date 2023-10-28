package usecase

import "context"

type Usecase interface {
	StartPipelinePeriodically(
		ctx context.Context,
	) error
	CrawlOnGCF(
		ctx context.Context,
		rawBytes []byte,
	) error
}
