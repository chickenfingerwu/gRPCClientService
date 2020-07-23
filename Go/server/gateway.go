package main

import (
	"context"
	"github.com/golang/glog"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"

	gw "./pbfile/service"
)

// newGateway returns a new gateway server which translates HTTP into gRPC.
func newGateway(ctx context.Context, conn *grpc.ClientConn) (http.Handler, error) {

	mux := gwruntime.NewServeMux()

	for _, f := range []func(context.Context, *gwruntime.ServeMux, *grpc.ClientConn) error{
		gw.RegisterServerServiceHandler,
	} {
		if err := f(ctx, mux, conn); err != nil {
			return nil, err
		}
	}
	return mux, nil
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//Configure http server
	mux := http.NewServeMux()
	conn, err := grpc.Dial(gRPCServerEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server %v", err.Error())
		return err
	}

	//Register new gateway server to translate http to gRPC
	gw, err := newGateway(ctx, conn)
	if err != nil {
		return err
	}
	mux.Handle("/", gw)

	s := &http.Server{
		Addr:    ":8081",
		Handler: allowCORS(mux),
	}

	go func() {
		<-ctx.Done()
		glog.Infof("Shutting down the http server")
		if err := s.Shutdown(context.Background()); err != nil {
			glog.Errorf("Failed to shutdown http server: %v", err)
		}
	}()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return s.ListenAndServe()
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		h.ServeHTTP(w, r)
	})
}
