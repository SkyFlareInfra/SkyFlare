package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/SkyFlareInfra/SkyFlare/application"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(application.Module)

	err := app.Start(context.Background())
	if err != nil {
		panic(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	err = app.Stop(context.Background())
	if err != nil {
		panic(err)
	}
}
