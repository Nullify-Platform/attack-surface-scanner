package main

import (
	"net/http"
	"testserver/pkg/config"
	"testserver/pkg/middleware"
	"testserver/pkg/routes"

	"github.com/go-chi/chi/v5"
	"github.com/nullify-platform/logger/pkg/logger"
)

func main() {
	log, err := logger.ConfigureProductionLogger(config.GetLogLevel())
	if err != nil {
		logger.Fatal("error configuring logger", logger.Err(err))
	}
	defer log.Sync()

	r := chi.NewRouter()

	r.Use(middleware.LoggingMiddleware)

	r.Mount("/api/users", routes.UsersRoutes())

	srv := http.Server{
		Addr:    ":8888",
		Handler: r,
	}

	logger.Info("starting server", logger.String("addr", srv.Addr))

	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
