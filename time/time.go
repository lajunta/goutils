package time

import (
	"time"
)

// PageNum caculate the page number
func PageNum(page int64, num int) int64 {
	return (page-1)*10 + 1 + int64(num)
}

// ToDate return just data
func ToDate(tm time.Time) string {
	loc := time.FixedZone("UTC+8", +8*60*60)
	return tm.In(loc).Format("2006-01-02")
}

// ToDate1 return timestring
func ToDate1(tm time.Time) string {
	loc := time.FixedZone("UTC+8", +8*60*60)
	return tm.In(loc).Format("200601021504")
}

// ToDate2 return timestring with delimiter
func ToDate2(tm time.Time) string {
	loc := time.FixedZone("UTC+8", +8*60*60)
	return tm.In(loc).Format("2006-01-02 15:04")
}
