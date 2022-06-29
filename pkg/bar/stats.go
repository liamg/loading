package bar

import (
	"fmt"
	"math"
)

type StatsFunc func(current int64, total int64, fraction float64, h History) string

func StatsPercentComplete(_ int64, total int64, fraction float64, _ History) string {
	if total == 0 {
		return "???%"
	}
	return fmt.Sprintf("%3.0f%%", fraction*100)
}

func StatsAmountComplete(current int64, total int64, _ float64, _ History) string {
	totalStr := fmt.Sprintf("%d", total)
	maxWidth := len(fmt.Sprintf("%d", total))
	if currentWidth := len(fmt.Sprintf("%d", current)); currentWidth > maxWidth {
		maxWidth = currentWidth
	}
	return fmt.Sprintf(
		fmt.Sprintf("%%%dd/%%s", maxWidth),
		current,
		totalStr)
}

func StatsTimeRemaining(_ int64, total int64, _ float64, h History) string {
	if total == 0 {
		return "ETA: " + unknownETA
	}
	return fmt.Sprintf("ETA: %5s", h.ETA(total))
}

func StatsBytesComplete(current int64, total int64, _ float64, _ History) string {
	return fmt.Sprintf("%12s", fmt.Sprintf("%s/%s", formatBytesAsString(current), formatBytesAsString(total)))
}

func StatsAmountRate(_ int64, total int64, _ float64, h History) string {
	rateMS := h.AverageRatePerMS()
	maxWidth := len(fmt.Sprintf("%d", total))

	if rateS := int64(math.Round(rateMS * 1_000)); rateS > 0 {
		return fmt.Sprintf(fmt.Sprintf("%%%dd/s", maxWidth), rateS)
	}
	if rateM := int64(math.Round(rateMS * 1_000 * 60)); rateM > 0 {
		return fmt.Sprintf(fmt.Sprintf("%%%dd/m", maxWidth), rateM)
	}
	if rateH := int64(math.Round(rateMS * 1_000 * 60 * 60)); rateH > 0 {
		return fmt.Sprintf(fmt.Sprintf("%%%dd/h", maxWidth), rateH)
	}
	return fmt.Sprintf(fmt.Sprintf("%%%ds/s", maxWidth), "??")
}

func StatsBytesRate(_ int64, _ int64, _ float64, h History) string {
	rateMS := h.AverageRatePerMS()
	rateS := int64(math.Round(rateMS * 1_000))
	return fmt.Sprintf("%6s/s", formatBytesAsString(rateS))
}

const siUnits = "kMGTPEZY"

// we use SI, not IEC - see https://wiki.ubuntu.com/UnitsPolicy
func formatBytesAsString(count int64) string {
	if count < 1000 {
		return fmt.Sprintf("%dB", count)
	}
	var offset int
	var divider int64 = 1000
	for n := count / 1000; n >= 1000; n /= 1000 {
		divider *= 1000
		offset++
	}
	return fmt.Sprintf("%.1f%cB",
		float64(count)/float64(divider), siUnits[offset])
}
