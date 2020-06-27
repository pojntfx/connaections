package main

import (
	"flag"
	"log"
	"net"

	proto "github.com/pojntfx/connaections/pkg/proto/generated"
	"github.com/pojntfx/connaections/pkg/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	laddr := flag.String("laddr", ":2004", "Listen address")

	flag.Parse()

	list, err := net.Listen("tcp", *laddr)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	reflection.Register(srv)

	csvc := services.Connections{}
	proto.RegisterConnectionsServer(srv, &csvc)

	log.Printf("Starting server on port %v", *laddr)

	if err := srv.Serve(list); err != nil {
		log.Fatal(err)
	}
}
