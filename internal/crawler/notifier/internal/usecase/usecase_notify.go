package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *UsecaseImpl) NotifyOnGCF(
	ctx context.Context,
	fullPath string,
) error {
	t.L.Infof(ctx, "NotifyOnGCP %s", fullPath)
	notifiers, err := t.NotifierFactory.GetNotiferFromDocPathFirestore(ctx, fullPath)
	if err != nil {
		return terrors.Wrap(err)
	}
	for _, notifier := range notifiers {
		t.L.Infof(ctx, "Notifier %s", notifier.ID())
		timeSeriesData := notifier.NewEmptyTimeSeriesData(ctx)
		if err := t.Repository.GetTimeSeriesDataFromFullPathFirestore(
			ctx,
			fullPath,
			timeSeriesData,
		); err != nil {
			t.L.Errorf(ctx, "%+v on notifier %s", err, notifier.ID())
			continue
		}
		if err := notifier.Notify(ctx, timeSeriesData); err != nil {
			t.L.Errorf(ctx, "%+v on notifier %s", err, notifier.ID())
			continue
		}
	}
	return nil
}
