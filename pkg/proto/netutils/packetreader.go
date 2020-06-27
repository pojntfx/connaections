package netutils

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type PacketReader struct {
	Dev        string
	PacketChan chan gopacket.Packet

	packetSource  *gopacket.PacketSource
	parser        *gopacket.DecodingLayerParser
	decodedLayers []gopacket.LayerType
}

func (pr *PacketReader) Open() error {
	handle, err := pcap.OpenLive(pr.Dev, 1600, true, pcap.BlockForever)
	if err != nil {
		return err
	}

	pr.packetSource = gopacket.NewPacketSource(handle, handle.LinkType())

	pr.decodedLayers = make([]gopacket.LayerType, 0, 10)

	var ipv4 layers.IPv4
	var ipv6 layers.IPv6
	var tcp layers.TCP
	var udp layers.UDP

	pr.parser = gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &ipv4, &ipv6, &tcp, &udp)

	return nil
}
