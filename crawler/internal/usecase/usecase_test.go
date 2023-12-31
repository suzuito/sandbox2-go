package usecase

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/queue"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
	"go.uber.org/mock/gomock"
)

type utMocks struct {
	MockLogBuffer                      *bytes.Buffer
	MockCrawlerFactory                 *factory.MockCrawlerFactory
	MockCrawlerRepository              *repository.MockCrawlerRepository
	MockNotifierRepository             *repository.MockNotifierRepository
	MockNotifierFactory                *factory.MockNotifierFactory
	MockCrawlerConfigurationRepository *repository.MockCrawlerConfigurationRepository
	MockTriggerCrawlerQueue            *queue.MockTriggerCrawlerQueue
	MockTimeSeriesDataRepository       *repository.MockTimeSeriesDataRepository

	MockFetcher   *crawler.MockFetcher
	MockParser    *crawler.MockParser
	MockPublisher *crawler.MockPublisher
	MockNotifier  *notifier.MockNotifier
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
		L:                              l,
		CrawlerRepository:              t.MockCrawlerRepository,
		NotifierRepository:             t.MockNotifierRepository,
		CrawlerFactory:                 t.MockCrawlerFactory,
		NotifierFactory:                t.MockNotifierFactory,
		TimeSeriesDataRepository:       t.MockTimeSeriesDataRepository,
		CrawlerConfigurationRepository: t.MockCrawlerConfigurationRepository,
		TriggerCrawlerQueue:            t.MockTriggerCrawlerQueue,
	}
}

func newUTMocks(ctrl *gomock.Controller) *utMocks {
	return &utMocks{
		MockLogBuffer:                      bytes.NewBufferString(""),
		MockCrawlerFactory:                 factory.NewMockCrawlerFactory(ctrl),
		MockNotifierRepository:             repository.NewMockNotifierRepository(ctrl),
		MockNotifierFactory:                factory.NewMockNotifierFactory(ctrl),
		MockCrawlerRepository:              repository.NewMockCrawlerRepository(ctrl),
		MockCrawlerConfigurationRepository: repository.NewMockCrawlerConfigurationRepository(ctrl),
		MockTriggerCrawlerQueue:            queue.NewMockTriggerCrawlerQueue(ctrl),
		MockTimeSeriesDataRepository:       repository.NewMockTimeSeriesDataRepository(ctrl),
		MockFetcher:                        crawler.NewMockFetcher(ctrl),
		MockParser:                         crawler.NewMockParser(ctrl),
		MockPublisher:                      crawler.NewMockPublisher(ctrl),
		MockNotifier:                       notifier.NewMockNotifier(ctrl),
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
