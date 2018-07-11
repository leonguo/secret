package main

import (
	"./app"
	"./config"
	"os"
	"os/signal"
	"time"
	"context"
)

func main() {
	// Start server
	go func() {
		if err := app.Server.Start(config.AppConfig.GetString("system.port")); err != nil {
			app.Server.Logger.Info("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Server.Shutdown(ctx); err != nil {
		app.Server.Logger.Fatal(err)
	}
}
