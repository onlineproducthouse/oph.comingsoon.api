package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"oph.api.comingsoon/oph.api.comingsoon.app/app"
	infrastructure "oph.api.comingsoon/oph.api.comingsoon.infrastructure"
	ioc "oph.api.comingsoon/oph.api.comingsoon.ioc"
)

func main() {
	infra := infrastructure.NewInfrastructure()
	iocContainer := ioc.NewIoc(infra)
	_app := app.NewApi(infra, iocContainer).Build()

	run(infra, iocContainer, _app)
}

func run(infra infrastructure.IInfrastructure, ioc ioc.IIoc, api app.IApi) {
	// Start server
	go func() {
		if err := api.Start(); err != nil && err != http.ErrServerClosed {
			infra.Logger().Fatal(fmt.Sprintf("shutting down the server: %v", err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := api.Stop(ctx); err != nil {
		infra.Logger().Fatal(err.Error())
	}
}
