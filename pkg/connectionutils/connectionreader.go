package connectionutils

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

type ConnectionReader struct {
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

func (cr *ConnectionReader) Open() error {
	handle, err := pcap.OpenLive(cr.Dev, 1600, true, pcap.BlockForever)
	if err != nil {
		return err
	}

	if err := handle.SetBPFFilter("tcp or udp"); err != nil {
		return err
	}

	cr.packetSource = gopacket.NewPacketSource(handle, handle.LinkType())

	cr.decodedLayers = make([]gopacket.LayerType, 0, 10)

	var eth layers.Ethernet

	cr.parser = gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &cr.ipv4, &cr.ipv6, &cr.tcp, &cr.udp)
	cr.parser.IgnoreUnsupported = true

	return nil
}

func (cr *ConnectionReader) Read() error {
	for packet := range cr.packetSource.Packets() {
		if err := cr.parser.DecodeLayers(packet.Data(), &cr.decodedLayers); err != nil {
			return err
		}

		connection := Connection{}

		for _, layerType := range cr.decodedLayers {
			switch layerType {
			case layers.LayerTypeIPv4:
				connection.SrcIP = cr.ipv4.SrcIP.String()
				connection.DstIP = cr.ipv4.SrcIP.String()
			case layers.LayerTypeIPv6:
				connection.SrcIP = cr.ipv6.SrcIP.String()
				connection.DstIP = cr.ipv6.SrcIP.String()

			case layers.LayerTypeTCP:
				connection.SrcPort = cr.tcp.SrcPort.String()
				connection.DstPort = cr.tcp.DstPort.String()
			case layers.LayerTypeUDP:
				connection.SrcPort = cr.udp.SrcPort.String()
				connection.DstPort = cr.udp.DstPort.String()
			}
		}

		cr.ConnectionChan <- connection
	}

	return nil
}
