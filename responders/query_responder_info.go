package responders

import (
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
	additionalQName           string
}
