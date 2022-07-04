package responders

import (
	"github.com/galenliu/dnssd/core"
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
	additionalQName           core.FullQName
}
