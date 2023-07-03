package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
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
			fmt.Println("Adding route: ", handler.Method(), handler.Pattern())
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

func Serve() {
	ServerDependencies := fx.Provide(
		NewHTTPClient,
		NewHTTPRouter,
	)

	app := fx.New(
		fx.Options(
			ServerDependencies,
			CalculateRideModule,
		),
		fx.Invoke(StartHTTPServer, NewLogger),
	)

	app.Run()
}
