package services

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pojntfx/connaections/pkg/packetutils"
	proto "github.com/pojntfx/connaections/pkg/proto/generated"
)

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"

type Connections struct {
	proto.UnimplementedConnectionsServer
	PacketReader *packetutils.PacketReader
}

func (c *Connections) Subscribe(req *empty.Empty, srv proto.Connections_SubscribeServer) error {
	for connection := range c.PacketReader.ConnectionChan {
		srv.Send(&proto.Connection{SrcIP: connection.SrcIP, SrcPort: connection.SrcPort, DstIP: connection.DstIP, DstPort: connection.DstPort})
	}

	return nil
}
