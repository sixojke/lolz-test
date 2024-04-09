package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/sixojke/lolz-test/internal/server"

	"github.com/sixojke/lolz-test/internal/config"
	"github.com/sixojke/lolz-test/internal/delivery"
	"github.com/sixojke/lolz-test/internal/repository"
	"github.com/sixojke/lolz-test/internal/service"
	"github.com/sixojke/lolz-test/pkg/migrations"
)

func Run() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("config error: %v", err))
	}

	postgres, err := repository.NewPostgresDB(cfg.Postgres)
	if err != nil {
		log.Fatal(fmt.Sprintf("postgres connection error: %v", err))
	}
	defer postgres.Close()
	log.Info("[POSTGRES] Connection successful")

	if err := migrations.MigratePostgres(cfg.Postgres); err != nil {
		log.Error(fmt.Sprintf("postgres migrate error: %v", err))
	}
	log.Info("[POSTGRES] Migrate successful")

	repo := repository.NewRepository(&repository.Deps{
		Postgres: postgres,
	})
	services := service.NewService(&service.Deps{
		Repo: repo,
	})
	handler := delivery.NewHandler(services, cfg.HandlerConfig)

	srv := server.NewServer(cfg.HTTPServer, handler.Init())
	go func() {
		if err := srv.Start(); !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error occurred while running http server: %v\n", err)
		}
	}()
	log.Info(fmt.Sprintf("[SERVER] Started :%v", cfg.HTTPServer.Port))

	shutdown(srv, postgres)
}

func shutdown(srv *server.Server, postgres *sqlx.DB) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 3 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Errorf("failed to stop server: %v", err)
	}

	postgres.Close()
}
