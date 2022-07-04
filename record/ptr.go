package record

import (
	"github.com/galenliu/dnssd/core"
	"github.com/galenliu/dnssd/core/QType"
	"github.com/miekg/dns"
)

type PtrResourceRecord struct {
	*Resource
	mPtrName core.FullQName
}

func NewPtrResourceRecord(qName, ptrName core.FullQName) *PtrResourceRecord {
	return &PtrResourceRecord{
		Resource: &Resource{
			mQType:      QType.PTR,
			mQname:      qName,
			mCacheFlush: false,
		},
		mPtrName: ptrName,
	}
}

func (r *PtrResourceRecord) GetPtr() core.FullQName {
	h := dns.RR_Header{}
	h.Header()
	return r.mPtrName
}
