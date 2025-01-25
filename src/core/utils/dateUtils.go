package utils

import (
	"strings"
	"time"
)

func ConvertTimeToYearMonthDayFormat(t time.Time) string {
	strd := t.String()
	d := strings.Split(strd, " ")
	return d[0]
}
