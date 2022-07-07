package responders

import (
	"github.com/galenliu/dnssd/QName"
	"github.com/galenliu/dnssd/record"
	"github.com/miekg/dns"
)

const kTxtDefaultTtl = 4500

type TxtResponder struct {
	dns.TXT
	mRecord *record.TxtResourceRecord
}

func NewTxtResponder(qname QName.FullQName, txt []string) *TxtResponder {
	return &TxtResponder{
		TXT: dns.TXT{
			Hdr: dns.RR_Header{
				Name:     qname.String(),
				Rrtype:   dns.TypeTXT,
				Class:    dns.ClassINET,
				Ttl:      kTxtDefaultTtl,
				Rdlength: 0,
			},
			Txt: txt,
		},
	}
}

func (txt *TxtResponder) GetName() string {
	return txt.TXT.Hdr.Name
}

func (txt *TxtResponder) GetType() uint16 {
	return txt.TXT.Hdr.Rrtype
}

func (txt *TxtResponder) GetClass() uint16 {
	return txt.TXT.Hdr.Class
}

func (txt *TxtResponder) GetTtl() uint32 {
	return txt.TXT.Hdr.Ttl
}

func (txt *TxtResponder) ResourceRecord() dns.RR {
	return &txt.TXT
}

func (txt *TxtResponder) SetTtl(ttl uint32) {
	txt.TXT.Hdr.Ttl = ttl
}
