package responders

import (
	"github.com/galenliu/dnssd/record"
	"github.com/miekg/dns"
)

type SrvResponder struct {
	dns.SRV
	mRecord *record.SrvResourceRecord
}

func NewSrvResponder(r *record.SrvResourceRecord) *SrvResponder {
	return &SrvResponder{
		SRV: dns.SRV{
			Hdr: dns.RR_Header{
				Name:     "",
				Rrtype:   dns.TypeSRV,
				Class:    dns.ClassINET,
				Ttl:      kDefaultTtl,
				Rdlength: 0,
			},
			Priority: 0,
			Weight:   0,
			Port:     0,
			Target:   "",
		},
		mRecord: r,
	}
}

func (srv SrvResponder) GetName() string {
	return srv.SRV.Hdr.Name
}

func (srv SrvResponder) GetType() uint16 {
	return srv.SRV.Hdr.Rrtype
}

func (srv SrvResponder) GetClass() uint16 {
	return srv.SRV.Hdr.Class
}

func (srv SrvResponder) GetTtl() uint32 {
	return srv.SRV.Hdr.Ttl
}

func (srv SrvResponder) SetTtl(ttl uint32) {
	srv.SRV.Hdr.Ttl = ttl
}

func (srv SrvResponder) ResourceRecord() dns.RR {
	return &srv.SRV
}
