package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xebia/go-exercise/internal/registration/datastore"
	"github.com/xebia/go-exercise/internal/registration/email"
	"github.com/xebia/go-exercise/pkg/slogx"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	datastoreService datastore.Service
	emailService     email.Service
	router           *http.ServeMux
	stopCh           chan os.Signal
}

func NewServer(ds datastore.Service, es email.Service) *Server {
	s := &Server{
		datastoreService: ds,
		emailService:     es,
		router:           http.NewServeMux(),
		stopCh:           make(chan os.Signal, 1),
	}

	s.router.HandleFunc("POST /patients/new", s.createPatient)
	return s
}

func (s *Server) ListenAndServe(addr string) {
	hs := &http.Server{
		Addr:              addr,
		Handler:           h2c.NewHandler(s.router, &http2.Server{}),
		ReadHeaderTimeout: 10 * time.Second,
		MaxHeaderBytes:    8 * 1024,
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		slogx.Fatal(err)
	}

	signal.Notify(s.stopCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		slog.Info(fmt.Sprintf("listening on %s", addr))

		err := hs.Serve(listener)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			slogx.Fatal(err)
		}
	}()

	<-s.stopCh
	slog.Info("shutting down the server...")

	// After SIGTERM Cloud Run waits maximum 10 seconds
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	if err := hs.Shutdown(ctxTimeout); err != nil {
		slog.Warn(fmt.Sprintf("server shutdown resulted in an error: %v", err))
	}
	slog.Info("server stopped")
}

func (s *Server) Shutdown() {
	s.stopCh <- os.Interrupt
}

func writeError(w http.ResponseWriter, errorMessage string, statusCode int) {
	slog.Error(errorMessage, "code", statusCode)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(errorMessage)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to respond with errorMessage: %s, error: %v", errorMessage, err))
	}
}
