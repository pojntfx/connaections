package services

import proto "github.com/pojntfx/connaections/pkg/proto/generated"

//go:generate sh -c "mkdir -p ../proto/generated && protoc --go_out=paths=source_relative,plugins=grpc:../proto/generated -I=../proto ../proto/*.proto"

type Connections struct {
	proto.UnimplementedConnectionsServer
}
