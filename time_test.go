package main

import (
	"testing"
	"time"
)

func TestDurationInWords(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		// minutes
		{"1m", "1 minute"},
		{"2m", "2 minutes"},
		{"59m59s", "59 minutes"},
		// hours
		{"60m", "1 hour"},
		{"1h0m1s", "1 hour"},
		{"2h", "2 hours"},
		// TODO, if less than 5 hours, include minutes
		{"23h59m59s", "23 hours"},
		// days
		{"24h", "1 day"},
		{"48h", "2 days"},
		{"72h", "3 days"},
		// TODO, if less than 3 days, include hours
	}

	for _, test := range tests {
		duration, _ := time.ParseDuration(test.input)
		got := timeDurationInWords(duration)
		if got != test.want {
			t.Errorf("timeDurationInWords(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestDurationDays(t *testing.T) {
	var tests = []struct {
		input string
		want  float64
	}{
		{"24h", 1},
		{"25h", 1.0416666666666667},
		{"48h", 2},
	}

	for _, test := range tests {
		duration, _ := time.ParseDuration(test.input)
		got := timeDurationDays(duration)
		if got != test.want {
			t.Errorf("timeDurationDays(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}
