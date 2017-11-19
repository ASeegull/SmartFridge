package agent

import (
	"flag"
	"net"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

const (
	DefaultAddr = "localhost:9001"
)

func Run() {
	flag.Parse()
	lis, err := net.Listen("tcp", DefaultAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	grpcServer.Serve(lis)
}
