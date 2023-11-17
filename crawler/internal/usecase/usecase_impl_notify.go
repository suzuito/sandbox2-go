package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *UsecaseImpl) NotifyOnGCF(
	ctx context.Context,
	fullPath string,
) error {
	t.L.InfoContext(ctx, "NotifyOnGCP", "fullPath", fullPath)
	notifierDefs, err := t.NotifierRepository.GetNotiferDefinitionsFromDocPathFirestore(ctx, fullPath)
	if err != nil {
		return terrors.Wrap(err)
	}
	for _, notifierDef := range notifierDefs {
		logger := t.L.With("notifierID", notifierDef.ID)
		logger.InfoContext(ctx, "Notifier")
		notifier, err := t.NotifierFactory.Get(ctx, &notifierDef)
		if err != nil {
			logger.ErrorContext(ctx, "Failed to NotifierFactory.Get", "err", err)
			continue
		}
		timeSeriesData := notifier.NewEmptyTimeSeriesData()
		if err := t.TimeSeriesDataRepository.GetTimeSeriesDataFromFullPathFirestore(
			ctx,
			fullPath,
			timeSeriesData,
		); err != nil {
			t.L.ErrorContext(ctx, "Failed to GetTimeSeriesDataFromFullPathFirestore", "err", err)
			continue
		}
		if err := notifier.Notify(ctx, timeSeriesData); err != nil {
			t.L.ErrorContext(ctx, "Failed to Notify", "err", err)
			continue
		}
	}
	return nil
}
