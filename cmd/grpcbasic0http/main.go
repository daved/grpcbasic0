package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/daved/grpcbasic0/pb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	var rcpAddr, port string
	flag.StringVar(&rcpAddr, "rcp", ":3323", "rcp addr (default: ':3323')")
	flag.StringVar(&port, "http", ":3343", "http port (default: '3343')")
	flag.Parse()

	ctx := context.Background()
	m := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, m, rcpAddr, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot register service handler: %s\n", err)
		os.Exit(1)
	}

	// custom routes first, and cors handling on all requests
	h := cors(preMuxRouter(m))

	if err = http.ListenAndServe(port, h); err != nil {
		fmt.Fprintf(os.Stderr, "http server error: %s\n", err)
		os.Exit(1)
	}
}

// cross-origin resource sharing
func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()

		// TODO: be careful!
		o := r.Header.Get("Origin")
		h.Set("Access-Control-Allow-Origin", o)

		if r.Method == http.MethodOptions {
			h.Set("Access-Control-Allow-Methods", strings.Join(
				[]string{
					http.MethodOptions,
					http.MethodGet,
					http.MethodPut,
					http.MethodHead,
					http.MethodPost,
					http.MethodDelete,
					http.MethodPatch,
					http.MethodTrace,
				}, ", ",
			))

			h.Set("Access-Control-Allow-Headers", strings.Join(
				[]string{
					"Access-Control-Allow-Headers",
					"Origin",
					"X-Requested-With",
					"Content-Type",
					"Accept",
				}, ", ",
			))

			return
		}

		next.ServeHTTP(w, r)
	})
}

// v1 swagger.json
func v1SwaggerHandler(w http.ResponseWriter, r *http.Request) {
	d, err := pb.Asset("grpcbasic0.swagger.json")
	if err == nil {
		_, err = w.Write(d)
		if err == nil {
			return
		}
	}

	fmt.Fprintf(os.Stderr, "cannot write swagger.json: %s\n", err)

	stts := http.StatusInternalServerError
	http.Error(w, http.StatusText(stts), stts)
}

// custom routing prior to generated api
func preMuxRouter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v1/openapi.json":
			http.HandlerFunc(v1SwaggerHandler).ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
