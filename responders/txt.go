package responders

import (
	"github.com/galenliu/dnssd/core/QType"
	"github.com/galenliu/dnssd/record"
)

type TxtResponder struct {
	*recordResponder
	mRecord *record.TxtResourceRecord
}

func NewTxtResponder(resourceRecord *record.TxtResourceRecord) *TxtResponder {
	return &TxtResponder{
		recordResponder: &recordResponder{
			responder: &responder{
				mQType: QType.SRV,
				mQName: resourceRecord.GetName(),
			},
			mTtl: kDefaultTtl,
		},
		mRecord: resourceRecord,
	}
}

func (r *TxtResponder) AddAllResponses(source *IPPacket.Info, delegate ResponderDelegate, configuration *ResponseConfiguration) {

}
