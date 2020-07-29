package main

import (
	"flag"
	"log"
	"net"
	"sync"

	proto "github.com/pojntfx/connaections/pkg/proto/generated"
	"github.com/pojntfx/connaections/pkg/readers"
	"github.com/pojntfx/connaections/pkg/services"
	"github.com/ugjka/messenger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	laddr := flag.String("laddr", ":2004", "Listen address")
	rdev := flag.String("rdev", "eth0", "Interface to read packets from")
	dbpath := flag.String("dbpath", "GeoLite2-City.mmdb", "Path to the GeoLite2 database")

	flag.Parse()

	cmsgr := messenger.New(0, false)
	cchan := make(chan readers.Connection)

	cr := readers.ConnectionReader{
		Dev:            *rdev,
		ConnectionChan: cchan,
	}

	if err := cr.Open(); err != nil {
		log.Fatal(err)
	}

	lr := readers.LocationReader{
		DbPath: *dbpath,
	}

	if err := lr.Open(); err != nil {
		log.Fatal(err)
	}

	cs := services.Connections{
		Messenger:      cmsgr,
		LocationReader: lr,
	}

	lis, err := net.Listen("tcp", *laddr)
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	reflection.Register(srv)

	proto.RegisterConnectionsServer(srv, &cs)

	log.Printf("Starting server on address %v", *laddr)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for connection := range cchan {
			cmsgr.Broadcast(connection)
		}

		wg.Done()
	}()

	go func() {
		if err := cr.Read(); err != nil {
			log.Fatal(err)
		}

		wg.Done()
	}()

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatal(err)
		}

		wg.Done()
	}()

	wg.Wait()
}
