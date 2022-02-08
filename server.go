package todo

import (
	"context"
	"net/http"
	"time"
)

type MyServer struct {
	httpServer *http.Server
}

func (server *MyServer) Run(port string, handler http.Handler) error {
	server.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server.httpServer.ListenAndServe()
}

func (server *MyServer) Shutdown(ctx context.Context) error {
	return server.httpServer.Shutdown(ctx)
}
