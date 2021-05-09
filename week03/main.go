package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {
	app := NewApp(Server(
		&http.Server{
			Addr: ":8080",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("hello errgroup"))
			}),
		},
		&http.Server{
			Addr: ":9090",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("hello signal"))
			}),
		},
	))

	err := app.Run(context.Background())
	if err != nil {
        fmt.Println(err)
    }
}

func NewApp(opts ...Opt) *App {
	var app App
	for _, opt := range opts {
		opt(&app.opts)
	}

	return &app
}

type options struct {
	servers []*http.Server
}

type Opt func(o *options)

func Server(srv ...*http.Server) Opt {
	return func(o *options) { o.servers = srv }
}

type App struct {
	opts options
}

func (a *App) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	g, ctx := errgroup.WithContext(ctx)
	for _, srv := range a.opts.servers {
		srv := srv
		g.Go(func() error {
			return srv.ListenAndServe()
		})
		g.Go(func() error {
			<-ctx.Done()
			err := srv.Shutdown(ctx)
			if err != nil {
				return errors.Wrap(ctx.Err(), err.Error())
			}
			return nil
		})
	}

	g.Go(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
		for {
			select {
			case <-ctx.Done():
				cancel()
				return ctx.Err()
			case s := <-c:
				cancel()
				fmt.Printf("elegent shutdown: [%v]\n", s)
			}
		}
	})

	err := g.Wait()
	if errors.Is(err, context.Canceled) {
		return nil
	}
	return err
}
