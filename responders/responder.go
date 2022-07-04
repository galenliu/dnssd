package responders

import (
	"github.com/miekg/dns"
)

type Responder interface {
	GetClass() uint16
	GetName() string
	GetType() uint16
	GetTtl() uint32
	SetTtl(uint32)
	ResourceRecord() dns.RR
}
