package main

import (
	"context"
	"os"

	"github.com/suzuito/sandbox2-go/internal/common/cusecase/clog"
)

func main() {
	ctx := context.Background()
	u, err := NewUsecaseLocal(ctx)
	if err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		os.Exit(1)
	}
	fullPath := "Crawler/TimeSeriesData/goconnpass/connpass-299108"
	// fullPath = "Crawler/TimeSeriesData/golangweekly/https---golangweekly.com-issues-476"
	if err := u.NotifyOnGCF(ctx, fullPath); err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		os.Exit(1)
	}
}
