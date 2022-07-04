package old

import "github.com/miekg/dns"

type ResourceRecords struct {
	dns.RR_Header
}

func (r ResourceRecords) GetName() string {
	return r.RR_Header.Name
}

func (r ResourceRecords) GetType() uint16 {
	return r.RR_Header.Rrtype
}

func (r ResourceRecords) GetClass() uint16 {
	return r.RR_Header.Class
}

func (r ResourceRecords) GetTtl() uint32 {
	return r.RR_Header.Ttl
}

func (r *ResourceRecords) SetName(n string) *ResourceRecords {
	r.RR_Header.Name = n
	return r
}

func (r *ResourceRecords) SetType(t uint16) *ResourceRecords {
	r.RR_Header.Rrtype = t
	return r
}

func (r *ResourceRecords) SetClass(c uint16) *ResourceRecords {
	r.RR_Header.Class = c
	return r
}

func (r *ResourceRecords) SetTtl(t uint32) *ResourceRecords {
	r.RR_Header.Ttl = t
	return r
}
