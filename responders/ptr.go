package responders

import (
	"github.com/galenliu/dnssd/QName"
	"github.com/miekg/dns"
)

type PtrResponder struct {
	dns.PTR
}

func NewPtrResponder(name, target QName.FullQName) *PtrResponder {
	ptr := &PtrResponder{
		PTR: dns.PTR{
			Hdr: dns.RR_Header{
				Name:     name.String(),
				Rrtype:   dns.TypePTR,
				Class:    dns.ClassINET,
				Ttl:      kDefaultTtl,
				Rdlength: 0,
			},
			Ptr: target.String(),
		},
	}
	return ptr
}

func (p PtrResponder) GetName() string {
	return p.PTR.Hdr.Name
}

func (p PtrResponder) GetType() uint16 {
	return p.PTR.Hdr.Rrtype
}

func (p PtrResponder) GetClass() uint16 {
	return p.PTR.Hdr.Class
}

func (p PtrResponder) GetTtl() uint32 {
	return p.PTR.Hdr.Ttl
}

func (p PtrResponder) SetTtl(ttl uint32) {
	p.PTR.Hdr.Ttl = ttl
}

func (p PtrResponder) ResourceRecord() dns.RR {
	return &p.PTR
}
