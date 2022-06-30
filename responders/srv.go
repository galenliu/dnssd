package responders

import (
	"github.com/galenliu/chip/inet/IPPacket"
	"github.com/galenliu/dnssd/core/QType"
	"github.com/galenliu/dnssd/record"
)

type SrvResponder struct {
	*recordResponder
	mRecord *record.SrvResourceRecord
}

func NewSrvResponder(resourceRecord *record.SrvResourceRecord) *SrvResponder {
	return &SrvResponder{
		recordResponder: &recordResponder{
			responder: &responder{
				mQType: QType.SRV,
				mQName: resourceRecord.GetName(),
			},
		},
		mRecord: resourceRecord,
	}
}

func (r *SrvResponder) AddAllResponses(source *IPPacket.Info, delegate ResponderDelegate, configuration *ResponseConfiguration) {

}
