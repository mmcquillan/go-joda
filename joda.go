package joda

import (
	"strings"
)

func Format(joda string) (goda string) {
	damap := map[string]string{
		"y":    "06",
		"Y":    "2006",
		"MMMM": "January",
		"MMM":  "Jan",
		"MM":   "01",
		"M":    "1",
		"EE":   "Tuesday",
		"E":    "Tue",
		"ee":   "02",
		"e":    "2",
		"hh":   "03",
		"h":    "3",
		"H":    "15",
		"mm":   "04",
		"m":    "4",
		"ss":   "05",
		"s":    "5",
		"S":    ".0",
		"a":    "PM",
		"z":    "MST",
		"Z":    "-0700",
	}
	goda = joda
	for j, g := range damap {
		goda = strings.Replace(goda, j, g, -1)
	}
	return goda
}
