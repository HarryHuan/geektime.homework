// 以下代码来源：https://blog.csdn.net/gxy_2016/article/details/116615568
// 作业对于刚入门Go的本人，有点费力，所以在借鉴，在学习。。。

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

// 启动 HTTP server
func StartHttpServer(srv *http.Server) error {
	http.HandleFunc("/hello", HelloServer)
	fmt.Println("http server start")
	err := srv.ListenAndServe()
	return err
}

// 增加一个 HTTP handler
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

// 1. 基于 errgroup 实现一个 http server 的启动和关闭，以及 liunx signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
func main() {
	// Background returns a non-nil, empty Context. It is never canceled, has no
	// values, and has no deadline. It is typically used by the main function,
	// initialization, and tests, and as the top-level Context for incoming
	// requests.
	ctx := context.Background()

	// 定义 withCancel -> cancel() 方法 去取消下游的 Context

	// WithCancel returns a copy of parent with a new Done channel. The returned
	// context's Done channel is closed when the returned cancel function is called
	// or when the parent context's Done channel is closed, whichever happens first.
	//
	// Canceling this context releases resources associated with it, so code should
	// call cancel as soon as the operations running in this Context complete.
	ctx, cancel := context.WithCancel(ctx)

	// 使用 errgroup 进行 goroutine 取消

	// WithContext returns a new Group and an associated Context derived from ctx.
	//
	// The derived Context is canceled the first time a function passed to Go
	// returns a non-nil error or the first time Wait returns, whichever occurs
	// first.
	group, errCtx := errgroup.WithContext(ctx)

	// http server
	srv := &http.Server{Addr: ":9090"}

	group.Go(func() error {
		return StartHttpServer(srv)
	})

	group.Go(func() error {
		// 阻塞。 因为 cancel、timeout、deadline 都可能导致 Done 被 close

		// Done returns a channel that's closed when work done on behalf of this
		// context should be canceled. Done may return nil if this context can
		// never be canceled. Successive calls to Done return the same value.
		// The close of the Done channel may happen asynchronously,
		// after the cancel function returns.
		<-errCtx.Done()

		fmt.Println("http server stop")

		// 关闭 http server
		return srv.Shutdown(errCtx)
	})

	// 这里要用 buffer 为 1 的 chan
	channel := make(chan os.Signal, 1)
	signal.Notify(channel)

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-channel:
				cancel()
			}
		}

		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("all group done!")
}
