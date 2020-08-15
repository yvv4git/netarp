package netarp

import (
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type ArpGenerator struct {
	iface   *net.Interface
	srcIpV4 net.IP
	dstIpV4 net.IP
}

// Задаем интерфейс
func (sender *ArpGenerator) SetIface(iface *net.Interface) {
	sender.iface = iface
}

// Задаем src ip
func (sender *ArpGenerator) SetSrcIp(ipAddres net.IP) {
	sender.srcIpV4 = ipAddres
}

// Возвращаем сетевой интерфейс
func (sender *ArpGenerator) GetIface() *net.Interface {
	return sender.iface
}

// Задаем dst ipv4 адрес из строки
func (sender *ArpGenerator) SetDstIpV4FromStr(ipv4 string) {
	dstIpAddr := net.ParseIP(ipv4)
	if ipv4 := dstIpAddr.To4(); ipv4 != nil {
		dstIpAddr = ipv4
	}

	sender.dstIpV4 = dstIpAddr
}

// Задаем dst ipv4 адрес
func (sender *ArpGenerator) SetDstIp(ipv4 net.IP) {
	sender.dstIpV4 = ipv4
}

// Создаемм arp пакет
func (sender *ArpGenerator) GenerateArpPackage() []byte {
	eth := layers.Ethernet{
		SrcMAC:       sender.iface.HardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		EthernetType: layers.EthernetTypeARP,
	}
	arp := layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPRequest,
		SourceHwAddress:   []byte(sender.iface.HardwareAddr),
		SourceProtAddress: []byte(sender.srcIpV4),
		DstHwAddress:      []byte{0, 0, 0, 0, 0, 0},
		DstProtAddress:    []byte(sender.dstIpV4),
	}

	// Настройка буфера и параметров для сериализации
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}

	// Собираем пакет
	gopacket.SerializeLayers(buf, opts, &eth, &arp)
	rawBytes := buf.Bytes()
	return rawBytes
}

// Конструктор
func NewArpSender() *ArpGenerator {
	return new(ArpGenerator)
}
