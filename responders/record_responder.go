package responders

import (
	"github.com/miekg/dns"
)

const kDefaultTtl uint32 = 120

type RecordResponder interface {
	Responder
	//AddAllResponses(info *IPPacket.Info, delegate ResponderDelegate, configuration *ResponseConfiguration)
}

type ResourceRecord interface {
	dns.RR
}
