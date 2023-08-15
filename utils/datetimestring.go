package utils

import (
	"strings"
	"time"
)

func DateTimeString() string {
	return strings.Replace(time.Now().UTC().Format("20060102150405.000"), ".", "", 1)
}
