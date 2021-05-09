package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()

	// 模拟相关的路由和请求
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "world")
	})

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	serverQuit := make(chan struct{})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverQuit <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	// 模拟服务启动
	g.Go(func() error {
		return server.ListenAndServe()
	})

	// 模拟服务errgroup的 ctx 退出
	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("errgroup exit...")
		case <-serverQuit:
			log.Println("server will out...")
		}

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		log.Println("shutting down server...")
		return server.Shutdown(ctx)
	})

	// 参考资料：
	// 1. [golang信号signal的处理](http://www.01happy.com/golang-signal/)
	// 2. [Go中的系统Signal处理 | Tony Bai](https://tonybai.com/2012/09/21/signal-handling-in-go/)
	// 3. [Golang的 signal - Go语言中文网 - Golang中文社区](https://studygolang.com/articles/2333)
	// 如何测试：
	// 1. MacOS 的iTerm2 中。ps aux | grep go |grep week03 找到go 进程，kill pid进程号即可(kill -9 pid 得不到对应效果）
	// 2. GoLand 的 Run 窗口中，Ctrl C 退出服务
	g.Go(func() error {
		c := make(chan os.Signal, 0)
		// 监听指定信号
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-c:
			return errors.Errorf("got the os signal: %v", sig)
		}
	})

	err := g.Wait()
	fmt.Println(err)
}
