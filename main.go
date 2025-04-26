package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/sdk/app"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context) error {
	return nil
}

func main() {
	app.Run(func(ctx context.Context, lg *zap.Logger, t *app.Telemetry) error {
		var arg struct {
			Addr string
		}
		flag.StringVar(&arg.Addr, "addr", ":8080", "address to listen on")
		flag.Parse()
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("Hello, world!"))
		})
		srv := &http.Server{
			Addr:    arg.Addr,
			Handler: mux,
		}
		g, ctx := errgroup.WithContext(ctx)
		g.Go(func() error {
			<-ctx.Done()
			return srv.Shutdown(ctx)
		})
		g.Go(func() error {
			err := srv.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				return errors.Wrap(err, "listen and serve")
			}
			return nil
		})

		return nil
	})
}
