package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type Server struct {
	Router *chi.Mux
	HTTP   *http.Server
	GRPC   *grpc.Server
}

func New() *Server {
	r := chi.NewRouter()

	return &Server{
		Router: r,
		HTTP: &http.Server{
			Addr:         ":3000",
			Handler:      r,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}
}

func (s *Server) SetupMiddleware() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
}

func (s *Server) StartHTTPServer() error {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	errs, ctx := errgroup.WithContext(ctx)

	log.Println("starting http server on port 3000")

	errs.Go(func() error {
		if err := s.HTTP.ListenAndServe(); err != nil {
			return fmt.Errorf("listen and serve error: %w", err)
		}
		return nil
	})

	<-ctx.Done()

	stop()
	log.Println("shutting down http server gracefully")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.HTTP.Shutdown(timeoutCtx); err != nil {
		log.Println(err.Error())
	}

	return nil
}

func (s *Server) StartGRPCServer() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	errs, ctx := errgroup.WithContext(ctx)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	log.Println("starting gRPC server on port 50051")

	s.GRPC = grpc.NewServer()

	errs.Go(func() error {
		if err := s.GRPC.Serve(listener); err != nil {
			return fmt.Errorf("gRPC serve error: %v", err)
		}
		return nil
	})

	<-ctx.Done()

	stop()
	log.Println("shutting down gRPC server gracefully")

	s.GRPC.GracefulStop()

	return nil
}
