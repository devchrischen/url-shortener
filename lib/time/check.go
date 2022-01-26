package time

import "time"

const YEAR float64 = 24 * 365

func CheckHashExpired(t time.Time) bool {
	hours := time.Since(t).Hours()
	if hours > YEAR {
		return true
	} else {
		return false
	}
}
