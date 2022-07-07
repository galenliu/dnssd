package responders

import (
	"github.com/galenliu/dnssd/QName"
	"time"
)

type QueryResponderRecord struct {
	reportService     bool
	LastMulticastTime time.Time
}

type QueryResponderInfo struct {
	QueryResponderRecord
	Responder                 Responder
	reportNowAsAdditional     bool
	alsoReportAdditionalQName bool
	additionalQName           QName.FullQName
}
