package mdns

import (
	"github.com/galenliu/gateway/pkg/log"
	"github.com/miekg/dns"
	"net/netip"
)

type QueryDelegate interface {
	OnQuery(r dns.ResponseWriter, data *QueryData)
}

type Server struct {
	servers               []*dns.Server
	mDelegate             QueryDelegate
	mIpv6BroadcastAddress netip.Addr
	mPort                 uint16
}

func (m Server) Init() *Server {
	m.servers = make([]*dns.Server, 0)
	return &m
}

func (m *Server) Shutdown() {
	for _, s := range m.servers {
		err := s.Shutdown()
		if err != nil {
			log.Info(err.Error())
		}
	}
}

func (m *Server) StartServer(adders []netip.Addr, port uint16) error {
	m.Shutdown()
	m.mPort = port
	var adder, net string
	if adders != nil {
		for _, a := range adders {
			if !a.IsValid() {
				log.Info("invalid addr")
				continue
			}
			if a.Is4() {
				net = "udp4"
				adder = netip.AddrPortFrom(a, port).String()
			}
			if a.Is6() {
				net = "udp6"
				adder = netip.AddrPortFrom(a, port).String()
			}
			m.servers = append(m.servers, &dns.Server{
				Addr:      adder,
				Net:       net,
				ReusePort: true,
			})
		}
	} else {
		if ad := netip.AddrPortFrom(netip.IPv4Unspecified(), port); ad.IsValid() {
			adder = ad.String()
			net = "udp4"
			m.servers = append(m.servers, &dns.Server{
				Addr:      adder,
				Net:       net,
				ReusePort: true,
			})
		}
		if ad := netip.AddrPortFrom(netip.IPv6Unspecified(), port); ad.IsValid() {
			adder = ad.String()

			net = "udp6"
			m.servers = append(m.servers, &dns.Server{
				Addr:      adder,
				Net:       net,
				ReusePort: true,
			})
		}
	}
	for _, s := range m.servers {
		s := s
		go func() {
			log.Infof("mDns listen  adder: %s", s.Addr)
			err := s.ListenAndServe()
			if err != nil {
				log.Infof("mDns server exit addr:%s", s.Addr)
				log.Info(err.Error())
			}
		}()
	}
	return nil
}

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

func (m *Server) Broadcast(msg *dns.Msg) error {
	clint := new(dns.Client)
	clint.Net = "udp"
	addr := netip.AddrPortFrom(GetIpv4Into(), m.mPort).String()
	_, _, err := clint.Exchange(msg, addr)
	if err != nil {
		log.Info(err.Error())
	}
	return nil
}

func GetIpv4Into() netip.Addr {
	addr := netip.AddrFrom4([4]byte{224, 0, 0, 251})
	return addr
}

func GetIpv6Into() netip.Addr {
	addr, _ := netip.ParseAddr("FF02::FB")
	return addr
}
