package main

import "fmt"

type VVT struct {
	Value     int
	ValueText string
}

type FgiResult struct {
	LastUpdated struct {
		EpochUnixSeconds int
		HumanDate        string
	}
	Fgi Fgi // XXX: Why just Fgi is not working?
}

type Fgi struct {
	Now           VVT
	PreviousClose VVT
	OneWeekAgo    VVT
	OneMonthAgo   VVT
	OneYearAgo    VVT
}

func valueToColor(v int) string {
	if 0 <= v && v <= 25 {
		return "🟢"
	} else if 26 <= v && v <= 50 {
		return "🟡"
	} else if 51 <= v && v <= 75 {
		return "🟠"
	} else { // it's over 75!
		return "🔴"
	}
}

func (v VVT) toString() string {
	color := valueToColor(v.Value)
	return fmt.Sprintf("%s %d (%s)", color, v.Value, v.ValueText)
}

func (fr FgiResult) toString() string {
	return fmt.Sprintf(`[lastUpdate: %s]
- now: %s
- prev: %s
- 1w ago: %s
- 1m ago: %s
- 1y ago: %s`, fr.LastUpdated.HumanDate,
		fr.Fgi.Now.toString(),
		fr.Fgi.PreviousClose.toString(),
		fr.Fgi.OneWeekAgo.toString(),
		fr.Fgi.OneMonthAgo.toString(),
		fr.Fgi.OneYearAgo.toString())
}
