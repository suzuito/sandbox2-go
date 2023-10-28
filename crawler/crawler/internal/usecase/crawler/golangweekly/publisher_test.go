package golangweekly

import (
	"context"
	"testing"

	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
	"go.uber.org/mock/gomock"
)

func TestPublish(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := repository.NewMockRepository(ctrl)
	repo.EXPECT().SetTimeSeriesData(gomock.Any(), gomock.Any())
	p := Publisher{
		repository: repo,
	}
	p.Publish(ctx)
}
