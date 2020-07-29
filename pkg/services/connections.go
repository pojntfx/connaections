package services

import (
	"github.com/golang/protobuf/ptypes/empty"
	proto "github.com/pojntfx/connaections/pkg/proto/generated"
	"github.com/pojntfx/connaections/pkg/readers"
	"github.com/ugjka/messenger"
)

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"

type Connections struct {
	proto.UnimplementedConnectionsServer

	Messenger      *messenger.Messenger
	LocationReader readers.LocationReader
}

func (c *Connections) Subscribe(req *empty.Empty, srv proto.Connections_SubscribeServer) error {
	client, err := c.Messenger.Sub()
	if err != nil {
		return err
	}

	for connection := range client {
		srcIP := connection.(readers.Connection).SrcIP
		dstIP := connection.(readers.Connection).DstIP

		srcLocation, err := c.LocationReader.GetLocationForIP(srcIP)
		if err != nil {
			return err
		}
		dstLocation, err := c.LocationReader.GetLocationForIP(dstIP)
		if err != nil {
			return err
		}

		srv.Send(&proto.Connection{
			SrcIP:          srcIP.String(),
			SrcPort:        connection.(readers.Connection).SrcPort,
			SrcLatitude:    srcLocation.Latitude,
			SrcLongitude:   srcLocation.Longitude,
			SrcCityName:    srcLocation.CityName,
			SrcCountryCode: srcLocation.CountryCode,
			DstIP:          dstIP.String(),
			DstPort:        connection.(readers.Connection).DstPort,
			DstLatitude:    dstLocation.Latitude,
			DstLongitude:   dstLocation.Longitude,
			DstCityName:    dstLocation.CityName,
			DstCountryCode: dstLocation.CountryCode,
		})
	}

	return nil
}
