package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type httpHandler struct{}

func (*httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

var g errgroup.Group

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", &httpHandler{})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// 系统信号终止
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	g.Go(func() error {
		<-done
		return server.Shutdown(context.Background())
	})

	g.Go(func() error{
		return server.ListenAndServe()
	})

	err := g.Wait()
	if err != nil {
		log.Print("Server closed")
	}
}