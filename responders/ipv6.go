package responders

import (
	"github.com/galenliu/dnssd/QName"
	"github.com/miekg/dns"
	"net"
)

type IPv6Responder struct {
	dns.AAAA
}

func NewIPv6Responder(qname QName.FullQName) *IPv6Responder {
	ip6 := &IPv6Responder{
		AAAA: dns.AAAA{
			Hdr: dns.RR_Header{
				Name:     qname.String(),
				Rrtype:   dns.TypeAAAA,
				Class:    dns.ClassINET,
				Ttl:      kDefaultTtl,
				Rdlength: 0,
			},
			AAAA: net.IP{},
		},
	}
	return ip6
}

func (ipv6 IPv6Responder) GetName() string {
	return ipv6.AAAA.Hdr.Name
}

func (ipv6 IPv6Responder) GetType() uint16 {
	return ipv6.AAAA.Hdr.Rrtype
}

func (ipv6 IPv6Responder) GetClass() uint16 {
	return ipv6.AAAA.Hdr.Class
}

func (ipv6 IPv6Responder) GetTtl() uint32 {
	return ipv6.AAAA.Hdr.Ttl
}

func (ipv4 IPv6Responder) SetTtl(ttl uint32) {
	ipv4.AAAA.Hdr.Ttl = ttl
}

func (ipv6 IPv6Responder) ResourceRecord() dns.RR {
	return &ipv6.AAAA
}
