package packetutils

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type Connection struct {
	SrcIP   string
	SrcPort string
	DstIP   string
	DstPort string
}

type PacketReader struct {
	Dev            string
	ConnectionChan chan Connection

	packetSource  *gopacket.PacketSource
	parser        *gopacket.DecodingLayerParser
	decodedLayers []gopacket.LayerType

	ipv4 layers.IPv4
	ipv6 layers.IPv6
	tcp  layers.TCP
	udp  layers.UDP
}

func (pr *PacketReader) Open() error {
	handle, err := pcap.OpenLive(pr.Dev, 1600, true, pcap.BlockForever)
	if err != nil {
		return err
	}

	if err := handle.SetBPFFilter("tcp and port 80"); err != nil {
		return err
	}

	pr.packetSource = gopacket.NewPacketSource(handle, handle.LinkType())

	pr.decodedLayers = make([]gopacket.LayerType, 0, 10)

	pr.parser = gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &pr.ipv4, &pr.ipv6, &pr.tcp, &pr.udp)

	return nil
}

func (pr *PacketReader) Read() error {
	for packet := range pr.packetSource.Packets() {
		if err := pr.parser.DecodeLayers(packet.Data(), &pr.decodedLayers); err != nil {
			return err
		}

		connection := Connection{}

		for _, layerType := range pr.decodedLayers {
			switch layerType {
			case layers.LayerTypeIPv4:
				connection.SrcIP = pr.ipv4.SrcIP.String()
				connection.DstIP = pr.ipv4.SrcIP.String()
			case layers.LayerTypeIPv6:
				connection.SrcIP = pr.ipv6.SrcIP.String()
				connection.DstIP = pr.ipv6.SrcIP.String()

			case layers.LayerTypeTCP:
				connection.SrcPort = pr.tcp.SrcPort.String()
				connection.DstPort = pr.tcp.DstPort.String()
			case layers.LayerTypeUDP:
				connection.SrcPort = pr.udp.SrcPort.String()
				connection.DstPort = pr.udp.DstPort.String()
			}
		}

		pr.ConnectionChan <- connection
	}

	return nil
}
