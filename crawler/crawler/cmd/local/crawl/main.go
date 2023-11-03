package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/crawler/crawler/cmd/local/internal"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
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
	crawlerInputDataString := flag.String("crawler-input-data", "{}", "CrawlerInputData")
	flag.Parse()
	if len(*crawlerIDString) <= 2 {
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
	// crawlerID := goconnpass.CrawlerID
	crawlerID := crawler.CrawlerID(*crawlerIDString)
	crawlerInputData := crawler.CrawlerInputData{}
	if err := json.Unmarshal([]byte(*crawlerInputDataString), &crawlerInputData); err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		os.Exit(1)
	}
	if err := u.Crawl(ctx, crawlerID, crawlerInputData); err != nil {
		clog.L.Errorf(ctx, "%+v", err)
		os.Exit(1)
	}
}
