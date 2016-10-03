package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/daved/grpcbasic0/idl"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func swaggerWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v1/swagger.json" {
			if d, err := idl.Asset("grpcbasic0.swagger.json"); err == nil {
				_, err = w.Write(d)
				if err != nil {
					fmt.Fprintf(os.Stderr, "cannot write swagger.json: %v\n", err)
				}

				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	var rcpAddr, port string
	flag.StringVar(&rcpAddr, "rcp", ":3323", "rcp addr (default: ':3323')")
	flag.StringVar(&port, "http", ":3343", "http port (default: '3343')")
	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	m := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := idl.RegisterUserServiceHandlerFromEndpoint(ctx, m, rcpAddr, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot register service handler: %v\n", err)
		os.Exit(1)
	}

	h := swaggerWrapper(m)

	if err = http.ListenAndServe(port, h); err != nil {
		fmt.Fprintf(os.Stderr, "http server error: %v\n", err)
		os.Exit(1)
	}
}
