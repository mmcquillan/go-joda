package joda

import (
	"testing"
	"time"
)

func TestConversions(t *testing.T) {
	check("E MMM e, H:mm", "Mon Jan 2, 15:04", t)
	check("Y-MM-ee H:mm:ssZ", "2006-01-02 15:04:05-0700", t)
}

func check(j string, g string, t *testing.T) {
	goda := Format(j)
	if goda != g {
		t.Fatalf("Failed to parse '" + j + "' => '" + g + "' from '" + goda + "'")
	}
	l, _ := time.LoadLocation("MST")
	f := time.Date(2006, time.January, 2, 15, 4, 5, 0, l).Format(goda)
	if g != f {
		t.Fatalf("Failed to format '" + g + "' => '" + f + "'")
	}
}
