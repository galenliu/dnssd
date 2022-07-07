package record

import (
	"github.com/galenliu/dnssd/QName"
	"github.com/galenliu/dnssd/QType"
	"github.com/miekg/dns"
)

type PtrResourceRecord struct {
	*Resource
	mPtrName QName.FullQName
}

func NewPtrResourceRecord(qName, ptrName QName.FullQName) *PtrResourceRecord {
	return &PtrResourceRecord{
		Resource: &Resource{
			mQType:      QType.PTR,
			mQname:      qName,
			mCacheFlush: false,
		},
		mPtrName: ptrName,
	}
}

func (r *PtrResourceRecord) GetPtr() QName.FullQName {
	h := dns.RR_Header{}
	h.Header()
	return r.mPtrName
}
