package server

import (
	"context"
	"fmt"
	"github.com/chjoaquim/ride-service/pkg/database"
	"github.com/go-chi/chi"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net"
	"net/http"
	"time"
)

type (
	HTTPHandler interface {
		Method() string
		Pattern() string
		http.Handler
	}

	RouterParams struct {
		fx.In
		Handlers []HTTPHandler `group:"handlers"`
	}
)

func NewLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}

func NewHTTPRouter(params RouterParams) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/ride", func(r chi.Router) {
		for _, handler := range params.Handlers {
			r.Method(handler.Method(), handler.Pattern(), handler)
		}
	})

	router.Route("/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
	})

	return router
}

func NewHTTPClient() *http.Client {
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second * 10,
			}).DialContext,
		},
	}

	return client
}

func StartHTTPServer(lc fx.Lifecycle, router *chi.Mux) {
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting the server...")
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping the server...")
			return srv.Shutdown(ctx)
		},
	})
}

func NewDatabase() *database.Database {
	// TODO: Change to configuration file.
	config := database.Config{
		Credential: database.Credential{
			Name:     "ride_db",
			Username: "postgres",
			Host:     "localhost",
			Schema:   "ride",
			Port:     5432,
			Password: "admin123",
		},
	}

	db := database.NewDatabase(&config)
	return db
}

func Serve() {
	ServerDependencies := fx.Provide(
		NewHTTPClient,
		NewHTTPRouter,
		NewDatabase,
	)

	app := fx.New(
		fx.Options(
			ServerDependencies,
			CalculateRideModule,
			PassengersModule,
			DriversModule,
		),
		fx.Invoke(StartHTTPServer, NewLogger),
	)

	app.Run()
}
