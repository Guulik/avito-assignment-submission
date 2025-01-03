package main

import (
	"Banner_Infrastructure/internal/app"
	"Banner_Infrastructure/internal/configure"
	sl "Banner_Infrastructure/internal/lib/logger/slog"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envStage = "stage"
	envProd  = "prod"
)

func main() {
	cfg := configure.MustLoadConfig()

	log := setupLogger(cfg.Env)

	a := app.New(log, cfg)

	ctx := context.Background()

	go func() {
		a.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	if err := a.Stop(ctx); err != nil {
		fmt.Println(fmt.Errorf("failed to gracefully stop app err=%s", err))
	}

	fmt.Println("Gracefully stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envStage:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := sl.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
