package height

import (
	"github.com/filecoin-project/go-state-types/abi"
	"time"
)

func ToTime(height abi.ChainEpoch) time.Time {
	subHeight := height - 1851120
	loc, _ := time.LoadLocation("Local")
	CriterionTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-05-30 00:00:00", loc)
	h, _ := time.ParseDuration("0.5m")
	return CriterionTime.Add(time.Duration(subHeight) * h)
}
