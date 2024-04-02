package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/blog2/pkg/environment"
	"github.com/suzuito/sandbox2-go/blog2/pkg/inject"
)

func main() {
	ctx := context.Background()
	env := environment.Environment{}
	if err := envconfig.Process("", &env); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	u, logger, err := inject.NewUsecaseImpl(ctx, &env)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(2)
	}
	if err := u.CreateTestData001(ctx); err != nil {
		logger.Error("", "err", err)
		os.Exit(3)
	}
}
