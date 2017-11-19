package agent

import (
	"net/http"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ASeegull/SmartFridge/agent/config"
	"github.com/ASeegull/SmartFridge/agent/proto"
	log "github.com/sirupsen/logrus"
)

func newGateway(ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	endpoint := config.GetServerAddr()
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err := proto.RegisteragentHandlerFromEndpoint(ctx, mux, endpoint, dialOpts)
	if err != nil {
		return nil, err
	}
	err = proto.RegistersettingsHandlerFromEndpoint(ctx, mux, endpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

// Run sets up gRPC server with reverse-proxy to allow use REST API
func Run(address string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	gw, err := newGateway(ctx, opts...)
	if err != nil {
		return err
	}
	mux.Handle("/", gw)

	addr := config.GetAddr()
	log.Infof("The server is running on port: %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}