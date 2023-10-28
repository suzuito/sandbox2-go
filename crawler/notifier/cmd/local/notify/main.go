package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/crawler/notifier/cmd/local/internal"
)

func usage() {
	// Headline of usage
	fmt.Fprintf(os.Stderr, "Run notify\n")
	// Print command line option list
	flag.PrintDefaults()
}

func main() {
	flag.CommandLine.SetOutput(os.Stderr)
	flag.Usage = usage
	fullPath := flag.String("full-path", "", "Full path")
	flag.Parse()
	if len(*fullPath) <= 2 {
		usage()
		os.Exit(2)
	}

	clog.L.SetLevel(clog.LevelDebug)

	ctx := context.Background()
	u, err := internal.NewUsecaseLocal(ctx)
	if err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		os.Exit(1)
	}
	// fullPath := "Crawler/TimeSeriesData/goconnpass/connpass-299108"
	// fullPath = "Crawler/TimeSeriesData/golangweekly/https---golangweekly.com-issues-476"
	if err := u.NotifyOnGCF(ctx, *fullPath); err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		os.Exit(1)
	}
}
