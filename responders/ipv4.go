package responders

import (
	"github.com/miekg/dns"
	"net/netip"
)

type IPv4Responder struct {
	dns.A
}

func NewIPv4Responder(qname string) *IPv4Responder {
	ip4 := &IPv4Responder{
		A: dns.A{
			Hdr: dns.RR_Header{
				Name:     qname,
				Rrtype:   dns.TypeA,
				Class:    dns.ClassINET,
				Ttl:      0,
				Rdlength: 0,
			},
			A: netip.IPv4Unspecified().AsSlice(),
		},
	}
	return ip4
}

func (ipv4 IPv4Responder) GetName() string {
	return ipv4.A.Hdr.Name
}

func (ipv4 IPv4Responder) GetType() uint16 {
	return ipv4.A.Hdr.Rrtype
}

func (ipv4 IPv4Responder) GetClass() uint16 {
	return ipv4.A.Hdr.Class
}

func (ipv4 IPv4Responder) GetTtl() uint32 {
	return ipv4.A.Hdr.Ttl
}

func (ipv4 IPv4Responder) SetTtl(ttl uint32) {
	ipv4.A.Hdr.Ttl = ttl
}

func (ipv4 *IPv4Responder) ResourceRecord() dns.RR {
	return &ipv4.A
}
