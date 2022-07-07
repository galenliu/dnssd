package record

import (
	"github.com/galenliu/dnssd/QName"
	"github.com/galenliu/dnssd/QType"
)

type SrvResourceRecord struct {
	Resource
	mServerName QName.FullQName
	mPort       uint16
	mPriority   uint16
	mWeight     uint16
}

func NewSrvResourceRecord(qName QName.FullQName, serverName QName.FullQName, port uint16) *SrvResourceRecord {
	return &SrvResourceRecord{
		Resource: Resource{
			mTtl:        kDefaultTtl,
			mQType:      QType.SRV,
			mQname:      qName,
			mCacheFlush: false,
		},
		mServerName: serverName,
		mPort:       port,
		mPriority:   0,
		mWeight:     0,
	}
}
