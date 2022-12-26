package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"si-test/pkg/common"
	"si-test/pkg/config"
	deliv "si-test/services/source_product/delivery/http"
	source_product "si-test/services/source_product/repository"
	uc "si-test/services/source_product/usecase"
	"syscall"
)

func main() {
	logger := config.Logger()
	config.Environment()
	logger.Info("initializing service source product")

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	address := host + ":" + port

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		// we need to reserve to buffer size 1, so the notifier are not blocked
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
	}()

	db := common.BaseCapsule{
		Database: config.PostgreSql(ctx, logger),
	}

	repo := source_product.NewSourceProductRepository(db, logger)
	ucase := uc.NewSourceProductUseCase(repo)

	r := mux.NewRouter()
	initHandler(r, ucase)
	http.Handle("/", r)

	httpServer := &http.Server{
		Addr:    address,
		Handler: r,
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		logger.Info(fmt.Sprintf("starting %s service on %s", os.Getenv("SERVICE_NAME"),
			host+":"+port))
		return httpServer.ListenAndServe()
	})

	g.Go(func() error {
		<-gCtx.Done()
		log.Printf("got signal %v, attempting graceful shutdown", gCtx.Err())
		return httpServer.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		logger.Error("exit reason: ", zap.Error(err))
	}
}

func initHandler(r *mux.Router, usecase uc.SourceProductUseCase) {
	deliv.NewDeliveryHttpArea(r, usecase)
}
