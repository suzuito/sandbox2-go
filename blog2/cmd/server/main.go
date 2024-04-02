package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/blog2/internal/web"
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
	w := inject.NewWebImpl(ctx, &env, u, logger)
	e := gin.New()
	e.Use(gin.Recovery())
	web.SetRouter(e, w)
	e.Run()
}
