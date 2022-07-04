package mdns

import (
	"github.com/galenliu/chip/inet/IP"
	"github.com/galenliu/gateway/pkg/log"
	"github.com/miekg/dns"
	"net"
	"net/netip"
	"strconv"
	"sync"
)

type QueryDelegate interface {
	OnQuery(r dns.ResponseWriter, data *QueryData)
}

type Server struct {
	servers []*dns.Server
	//mResponseDelegate     PacketDelegate
	//mQueryDelegate        PacketDelegate
	mDelegate             QueryDelegate
	mIpv6BroadcastAddress netip.Addr
}

func (m *Server) ShutdownEndpoint(info EndpointInfo) {
	//TODO implement me
	panic("implement me")
}

var insServer *Server
var serOnce sync.Once

func GlobalServer() *Server {
	serOnce.Do(func() {
		insServer = newMdnsServer()
	})
	return insServer
}

func DefaultServer() *Server {
	serOnce.Do(func() {
		insServer = newMdnsServer()
	})
	return insServer
}

func newMdnsServer() *Server {
	return &Server{}
}

func (m *Server) Shutdown() {
	for _, s := range m.servers {
		err := s.Shutdown()
		if err != nil {
			log.Info(err.Error())
		}
	}
}

func (m *Server) StartServer(listeners []net.Listener, port uint16) error {
	m.Shutdown()
	if listeners != nil {
		for _, l := range listeners {
			server := &dns.Server{
				Net:               "",
				Listener:          l,
				TLSConfig:         nil,
				PacketConn:        nil,
				Handler:           m,
				UDPSize:           0,
				ReadTimeout:       0,
				WriteTimeout:      0,
				IdleTimeout:       nil,
				TsigProvider:      nil,
				TsigSecret:        nil,
				NotifyStartedFunc: nil,
				DecorateReader:    nil,
				DecorateWriter:    nil,
				MaxTCPQueries:     0,
				ReusePort:         true,
				MsgAcceptFunc:     nil,
			}
			m.servers = append(m.servers, server)
		}
		return nil
	}

	m.servers = append(m.servers, &dns.Server{
		Addr:              ":" + strconv.Itoa(int(port)),
		Net:               "udp",
		TLSConfig:         nil,
		PacketConn:        nil,
		Handler:           m,
		UDPSize:           0,
		ReadTimeout:       0,
		WriteTimeout:      0,
		IdleTimeout:       nil,
		TsigProvider:      nil,
		TsigSecret:        nil,
		NotifyStartedFunc: nil,
		DecorateReader:    nil,
		DecorateWriter:    nil,
		MaxTCPQueries:     0,
		ReusePort:         true,
		MsgAcceptFunc:     nil,
	})
	for _, s := range m.servers {
		s := s
		go func() {
			err := s.ListenAndServe()
			if err != nil {
				log.Info(err.Error())
			}
		}()
	}
	return nil
}

//func (m *Server) OnQuery(data *core.BytesRange, info *IPPacket.Info) {
//	if m.mQueryDelegate != nil {
//		m.mResponseDelegate.OnMdnsPacketData(data, info)
//	}
//}

func (m *Server) SetQueryDelegate(delegate QueryDelegate) {
	m.mDelegate = delegate
}

func (m *Server) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	m.mDelegate.OnQuery(w, &QueryData{
		Msg:                  r,
		mAnswerViaUnicast:    false,
		mIsInternalBroadcast: false,
	})
}

func GetIpv4Into() IP.Address {
	addr := netip.AddrFrom4([4]byte{224, 0, 0, 251})
	return IP.Address{Addr: addr}
}

func GetIpv6Into() IP.Address {
	addr, _ := netip.ParseAddr("FF02::FB")
	return IP.Address{Addr: addr}
}
