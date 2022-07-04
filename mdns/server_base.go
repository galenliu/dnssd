package mdns

import (
	"github.com/galenliu/chip/inet/IP"
	"github.com/galenliu/chip/inet/Interface"
	"github.com/galenliu/chip/inet/udp_endpoint"
	"github.com/galenliu/gateway/pkg/system"
	"net/netip"
)

//type ServerBase interface {
//	Shutdown()
//	SetDelegate()
//	ShutdownEndpoint(info EndpointInfo)
//	Listen(manager udp_endpoint.UDPEndpoint, port int) error
//	DirectSend(packet *system.PacketBufferHandle, address IP.Address, port int, id Interface.Id) error
//	BroadcastSend(packet *system.PacketBufferHandle, port int, id Interface.Id, addr IP.Address) error
//}

type DnsServer interface {
	Shutdown()
	SetDelegate()
	ShutdownEndpoint(info EndpointInfo)
	StartServer(port uint16)
	DirectSend(packet *system.PacketBufferHandle, address IP.Address, port int, id Interface.Id) error
	BroadcastSend(packet *system.PacketBufferHandle, port int, id Interface.Id, addr IP.Address) error
}

type BroadcastSendDelegate interface {
	Accept(info EndpointInfo) *udp_endpoint.UDPEndpoint
}

type ListenSocketPickerDelegate struct {
	BroadcastSendDelegate
}

//type PacketDelegate interface {
//	OnMdnsPacketData(data *core.BytesRange, info *IPPacket.Info)
//}

type InterfaceTypeFilterDelegate struct {
	BroadcastSendDelegate
	mInterface Interface.Id
	mAddress   netip.Addr
	mChild     BroadcastSendDelegate
}

func (d InterfaceTypeFilterDelegate) Accept(info EndpointInfo) *udp_endpoint.UDPEndpoint {
	return d.mChild.Accept(info)
}
