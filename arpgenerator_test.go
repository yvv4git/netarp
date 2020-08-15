package netarp

import (
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yvv4git/netipv4"
)

var ifaceName string
var dstIpAddrStr string

func setUp() {
	ifaceName = os.Getenv("TEST_IFACE")
	dstIpAddrStr = os.Getenv("TEST_DST_IP")
}

func TestGenerateArpPackage(t *testing.T) {
	setUp()
	t.Log(ifaceName)
	t.Log(dstIpAddrStr)

	// get network interface
	iface, err := net.InterfaceByName(ifaceName)
	if nil != err {
		panic("Don't find interface.")
	}

	// get src from network interface
	srcIp := netipv4.GetIpv4FromIface(iface)
	dstIp := net.ParseIP(dstIpAddrStr).To4()

	// create arp package
	arpSender := NewArpSender()
	arpSender.SetIface(iface)
	arpSender.SetSrcIp(srcIp)
	arpSender.SetDstIp(dstIp)
	//arpSender.SetDstIpV4FromStr(dstIpAddrStr)
	arpPackage := arpSender.GenerateArpPackage()

	// check
	t.Log(arpPackage)
	assert.NotEmpty(t, arpPackage)
	/* assert.Equal(t, []byte{255, 255, 255, 255, 255, 255, 0, 34,
		251, 196, 246, 212, 8, 6, 0, 1, 8, 0, 6, 4, 0, 1, 0, 34,
		251, 196, 246, 212, 192, 168, 1, 39, 0, 0, 0,
		0, 0, 0, 192, 168, 1, 1, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		arpPackage,
	) */
}
