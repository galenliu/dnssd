package responders

import (
	"github.com/galenliu/dnssd/core"
	"github.com/galenliu/dnssd/core/QClass"
	"github.com/galenliu/dnssd/core/QType"
	"github.com/galenliu/dnssd/record"
)

type ResponderDelegate interface {
	AddResponse(record record.ResourceRecord)
}

type responder struct {
	mQType QType.T
	mQName *core.FullQName
}

func newResponder(qType QType.T, mQname *core.FullQName) *responder {
	return &responder{
		mQType: qType,
		mQName: mQname,
	}
}

func (r responder) GetQClass() QClass.T {
	return QClass.IN
}

func (r responder) GetQType() QType.T {
	return r.mQType
}

func (r responder) GetQName() *core.FullQName {
	return r.mQName
}
