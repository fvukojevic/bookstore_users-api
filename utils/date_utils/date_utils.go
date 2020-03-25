package date_utils

import "time"

const (
	apiDayLayout = "2006-01-02T15:04:05Z"
	apiDbLayout  = "2006-01-02 15:04:02"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDayLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
