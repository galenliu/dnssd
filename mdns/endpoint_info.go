package mdns

import (
	"github.com/galenliu/chip/inet/udp_endpoint"
	"net/netip"
)

type EndpointInfo struct {
	mInterfaceId Interface.Id
	mAddress     netip.Addr
	mListenUdp   *udp_endpoint.UDPEndpoint
}
