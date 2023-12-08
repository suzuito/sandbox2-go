package usecase

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/queue"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
	"go.uber.org/mock/gomock"
)

type utMocks struct {
	MockLogBuffer           *bytes.Buffer
	MockCrawlerRepository   *repository.MockCrawlerRepository
	MockTriggerCrawlerQueue *queue.MockTriggerCrawlerQueue
}

func (t *utMocks) NewUsecase() *UsecaseImpl {
	h := slog.NewTextHandler(
		t.MockLogBuffer,
		&slog.HandlerOptions{
			AddSource: false,
			Level:     slog.LevelInfo,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				switch a.Key {
				case slog.SourceKey, slog.TimeKey:
					return slog.Attr{}
				}
				return a
			},
		},
	)
	l := slog.New(h)
	return &UsecaseImpl{
		L:                   l,
		CrawlerRepository:   t.MockCrawlerRepository,
		TriggerCrawlerQueue: t.MockTriggerCrawlerQueue,
	}
}

func newUTMocks(ctrl *gomock.Controller) *utMocks {
	return &utMocks{
		MockLogBuffer:           bytes.NewBufferString(""),
		MockCrawlerRepository:   repository.NewMockCrawlerRepository(ctrl),
		MockTriggerCrawlerQueue: queue.NewMockTriggerCrawlerQueue(ctrl),
	}
}

type utTestCase struct {
	desc             string
	setUp            func(mocks *utMocks)
	expectedLogLines []string
}

func (t *utTestCase) run(
	tt *testing.T,
	runner func(uc *UsecaseImpl),
) {
	tt.Run(t.desc, func(ttt *testing.T) {
		ctrl := gomock.NewController(tt)
		defer ctrl.Finish()
		mocks := newUTMocks(ctrl)
		t.setUp(mocks)
		uc := mocks.NewUsecase()
		runner(uc)
		assert.Equal(
			tt,
			strings.Join(t.expectedLogLines, "\n")+"\n",
			mocks.MockLogBuffer.String(),
		)
	})
}
