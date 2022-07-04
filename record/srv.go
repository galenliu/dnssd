package record

import (
	"github.com/galenliu/dnssd/core"
	"github.com/galenliu/dnssd/core/QType"
)

type SrvResourceRecord struct {
	Resource
	mServerName core.FullQName
	mPort       uint16
	mPriority   uint16
	mWeight     uint16
}

func NewSrvResourceRecord(qName core.FullQName, serverName core.FullQName, port uint16) *SrvResourceRecord {
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
