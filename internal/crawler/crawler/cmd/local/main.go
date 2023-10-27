package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/suzuito/sandbox2-go/internal/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/entity/crawler"
)

func usage() {
	// Headline of usage
	fmt.Fprintf(os.Stderr, "Run crawl\n")
	// Print command line option list
	flag.PrintDefaults()
}

func main() {
	flag.CommandLine.SetOutput(os.Stderr)
	flag.Usage = usage
	crawlerIDString := flag.String("crawler-id", "", "CrawlerID")
	flag.Parse()

	if len(*crawlerIDString) <= 2 {
		usage()
		os.Exit(2)
	}

	ctx := context.Background()
	u, err := NewUsecaseLocal(ctx)
	if err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		os.Exit(1)
	}
	// crawlerID := goconnpass.CrawlerID
	// crawlerID = golangweekly.CrawlerID
	crawlerID := crawler.CrawlerID(*crawlerIDString)
	if err := u.Crawl(ctx, crawlerID); err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		os.Exit(1)
	}
}
