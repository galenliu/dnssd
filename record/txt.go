package record

import (
	"github.com/galenliu/dnssd/core"
	"github.com/galenliu/dnssd/core/QType"
)

const kTxtDefaultTtl = 4500
const kMaxTxtRecordLength = 63

type TxtResourceRecord struct {
	Resource
	mEntries string
}

func NewTxtResourceRecord(qName core.FullQName) *TxtResourceRecord {
	return &TxtResourceRecord{
		Resource: Resource{
			mTtl:        kTxtDefaultTtl,
			mQType:      QType.TXT,
			mQname:      qName,
			mCacheFlush: false,
		},
	}
}
