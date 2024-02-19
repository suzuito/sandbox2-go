package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/inject"
	"github.com/suzuito/sandbox2-go/blog2/internal/web"
)

func main() {
	fmt.Println("TODO")
	ctx := context.Background()
	_, w, err := inject.NewImpl(ctx)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	e := gin.New()
	e.Use(gin.Recovery())
	web.SetRouter(e, w)
	e.Run()
}
