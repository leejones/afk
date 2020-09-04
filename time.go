package main

import (
	"fmt"
	"math"
	"time"
)

func timeDurationDays(d time.Duration) float64 {
	return d.Hours() / 24
}

func timeDurationInWords(d time.Duration) string {
	if d.Minutes() < 60 {
		minutes := math.Floor(d.Minutes())
		if minutes == 1 {
			return "1 minute"
		}
		return fmt.Sprintf("%v minutes", minutes)
	} else if d.Hours() < 24 {
		hours := math.Floor(d.Hours())
		if hours == 1 {
			return "1 hour"
		}
		return fmt.Sprintf("%v hours", hours)
	}
	days := math.Floor(timeDurationDays(d))
	if days == 1 {
		return "1 day"
	}
	return fmt.Sprintf("%v days", days)
}
