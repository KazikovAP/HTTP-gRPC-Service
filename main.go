package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"server/app/server"

	"golang.org/x/sync/errgroup"
)

var useHTTP, useGRPC bool

func init() {
	flag.BoolVar(&useHTTP, "http", false, "run HTTP server")
	flag.BoolVar(&useGRPC, "grpc", false, "run gRPC server")
}

func main() {
	fmt.Println("Welcome to Skeleton Project")

	flag.Parse()

	if !useHTTP && !useGRPC {
		log.Println("flag isn't specified")
		log.Println("running http server")
		useHTTP = true
	}

	s := server.New()
	s.SetupMiddleware()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errs, _ := errgroup.WithContext(ctx)

	if useHTTP {
		errs.Go(func() error {
			return s.StartHTTPServer()
		})
	}

	if useGRPC {
		errs.Go(func() error {
			return s.StartGRPCServer()
		})
	}

	if err := errs.Wait(); err != nil {
		log.Fatal(err)
	}

}
