package app

import (
	"context"
	"net/http"
	"time"
)

const (
	Bytes   = 20
	Timeout = 10
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << Bytes,
		ReadTimeout:    Timeout * time.Second,
		WriteTimeout:   Timeout * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
