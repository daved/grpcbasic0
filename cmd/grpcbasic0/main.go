package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/daved/grpcbasic0/pb"
	"google.golang.org/grpc"
)

func main() {
	var port string
	flag.StringVar(&port, "rcp", ":3323", "rcp port (default: ':3323')")
	flag.Parse()

	l, err := listener(port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot get listener: %s\n", err)
		os.Exit(1)
	}

	srvr, svc := grpc.NewServer(), NewUserService()
	pb.RegisterUserServiceServer(srvr, svc)

	if err = srvr.Serve(l); err != nil {
		fmt.Fprintf(os.Stderr, "rpc server error: %s\n", err)
		os.Exit(1)
	}
}

func listener(port string) (*net.TCPListener, error) {
	a, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve addr: %s", err)
	}

	l, err := net.ListenTCP("tcp", a)
	if err != nil {
		return nil, fmt.Errorf("cannot listen: %s", err)
	}

	return l, nil
}
