package time

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gotokatsuya/lboard/config"
)

func init() {
	location, err := time.LoadLocation("UTC")
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Panic("Time.LoadLocation")
	}
	time.Local = location
}

var appLocation *time.Location

func initAppLocation() {
	loc, err := time.LoadLocation(config.GetTime().Zone)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Panic("Time.LoadLocation")
	}
	appLocation = loc
}

const (
	dateTimeLayout = "2006-01-02 15:04:05"

	localDateTimeLayout          = "2006-01-02T15:04:05"
	localDateTimeLayoutWithoutSS = "2006-01-02T15:04"
)

// ParseInLocalDateTimeFormat ...
func ParseInLocalDateTimeFormat(value string) time.Time {
	initAppLocation()
	t, _ := time.ParseInLocation(dateTimeLayout, value, appLocation)
	if t.IsZero() {
		t, _ = time.ParseInLocation(localDateTimeLayout, value, appLocation)
	}
	if t.IsZero() {
		t, _ = time.ParseInLocation(localDateTimeLayoutWithoutSS, value, appLocation)
	}
	return t
}

// FormatInLocalDateTime ...
func FormatInLocalDateTime(date time.Time) string {
	return InLocal(date).Format(dateTimeLayout)
}

// Now returns current time
// 現在時刻を返却する
func Now() time.Time {
	return time.Now()
}

// DeltaMinutes adds minutes from now
// 現在時刻に指定の分を足した時刻を返却する
func DeltaMinutes(n time.Duration) time.Time {
	return Now().Add(n * time.Minute)
}

// InLocal ...
func InLocal(t time.Time) time.Time {
	initAppLocation()
	return t.In(appLocation)
}
