package services

import (
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pojntfx/connaections/pkg/connectionutils"
	proto "github.com/pojntfx/connaections/pkg/proto/generated"
	"github.com/ugjka/messenger"
)

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"

type Connections struct {
	proto.UnimplementedConnectionsServer

	Messenger     *messenger.Messenger
	CountryLookup connectionutils.CountryLookup
}

func (c *Connections) Subscribe(req *empty.Empty, srv proto.Connections_SubscribeServer) error {
	client, err := c.Messenger.Sub()
	if err != nil {
		return err
	}

	for connection := range client {
		srcIP := connection.(connectionutils.Connection).SrcIP
		dstIP := connection.(connectionutils.Connection).DstIP

		srv.Send(&proto.Connection{
			SrcIP:      srcIP,
			SrcPort:    connection.(connectionutils.Connection).SrcPort,
			SrcCountry: c.CountryLookup.GetCountryForIP(net.ParseIP(srcIP)),
			DstIP:      dstIP,
			DstPort:    connection.(connectionutils.Connection).DstPort,
			DstCountry: c.CountryLookup.GetCountryForIP(net.ParseIP(dstIP)),
		})
	}

	return nil
}
