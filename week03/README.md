# 问题
基于 `errgroup` 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

# 解法
## `errgroup`小结

**作用**
* 相比较于标准库的`waitgroup`，无需手动`Add`和`Done`
* 通过`WithContext`函数返回的`ctx`，`ctx.Done()`可以暗示`errgroup`中执行的任务是否完成或出错

**用法**
* `WithContext`
* `var g errgroup.Group`

## 代码
```golang
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
				fmt.Println("elegent shutdown")
				return ctx.Err()
			case <-c:
				cancel()
				fmt.Println("elegent shutdown")
			}
		}
	})

	err := g.Wait()
	if errors.Is(err, context.Canceled) {
		return nil
	}
	return err
}

```
