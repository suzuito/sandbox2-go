package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/suzuito/sandbox2-go/internal/blog/cmd"
	"github.com/suzuito/sandbox2-go/internal/blog/cmd/internal/inject"
)

func main() {
	ctx := context.Background()
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "subcommand is not found\n")
		os.Exit(1)
	}
	var subCmd cmd.SubCmd = nil
	switch os.Args[1] {
	case "convert":
		subCmd = &convertCmd{}
	default:
		fmt.Fprintf(os.Stderr, "subcommand '%s' is not found\n", os.Args[1])
		os.Exit(1)
	}
	if err := subCmd.Run(ctx, os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

type convertCmd struct {
}

func (t *convertCmd) Run(
	ctx context.Context,
	args []string,
) error {
	convertCmd := flag.NewFlagSet("convert", flag.ExitOnError)
	if err := convertCmd.Parse(args); err != nil {
		return err
	}
	return t.runLocalSources(ctx)
}

func (t *convertCmd) runLocalSources(
	ctx context.Context,
) error {
	u, _, _, err := inject.NewUsecaseImpl(ctx)
	if err != nil {
		return err
	}
	return u.UploadAllArticles(ctx, "main")
}
