package main

import (
	"flag"
	"log"
	"net"

	"github.com/pojntfx/connaections/pkg/packetutils"
	proto "github.com/pojntfx/connaections/pkg/proto/generated"
	"github.com/pojntfx/connaections/pkg/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	laddr := flag.String("laddr", ":2004", "Listen address")
	rdev := flag.String("rdev", "eth0", "Interface to read packets from")

	flag.Parse()

	list, err := net.Listen("tcp", *laddr)
	if err != nil {
		log.Fatal(err)
	}

	pr := packetutils.PacketReader{
		Dev:            *rdev,
		ConnectionChan: make(chan packetutils.Connection),
	}
	if err := pr.Open(); err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := pr.Read(); err != nil {
			log.Fatal(err)
		}
	}()

	srv := grpc.NewServer()
	reflection.Register(srv)

	csvc := services.Connections{
		PacketReader: &pr,
	}
	proto.RegisterConnectionsServer(srv, &csvc)

	log.Printf("Starting server on port %v", *laddr)

	if err := srv.Serve(list); err != nil {
		log.Fatal(err)
	}
}
