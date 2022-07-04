package record

import (
	"github.com/galenliu/dnssd/core"
	"github.com/galenliu/dnssd/core/QType"
)

const kDefaultTtl = 120

type Resource struct {
	mTtl        uint32
	mQType      QType.T
	mQname      core.FullQName
	mCacheFlush bool
}

func NewResourceRecord() *Resource {
	return &Resource{
		mTtl:        kDefaultTtl,
		mCacheFlush: false,
	}
}

func (r *Resource) SetTtl(u uint32) {
	r.mTtl = u
}

func (r *Resource) setCacheFlush(set bool) {
	r.mCacheFlush = set
}

func (r *Resource) getCacheFlush() bool {
	return r.mCacheFlush
}

func (r *Resource) getTtl() uint32 {
	return r.mTtl
}

func (r *Resource) GetName() core.FullQName {
	return r.mQname
}
