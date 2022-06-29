package bar

import (
	"fmt"
	"time"
)

type History []Record

const historyBufferSize = 128

const unknownETA = "??m??s"

func (h History) ETA(total int64) string {
	avgRate := h.AverageRatePerMS()
	if avgRate == 0 {
		return unknownETA
	}
	ms := int64(float64(total-h[len(h)-1].Progress) / avgRate)
	duration := time.Duration(ms * 1_000_000)

	if duration.Hours() >= 1 {
		return fmt.Sprintf("%02dh%02dm", int(duration.Hours()), int(duration.Minutes())%60)
	}

	return fmt.Sprintf("%02dm%02ds", int(duration.Minutes()), int(duration.Seconds())%60)
}

func (h *History) Push(record Record) {
	newH := append(*h, record)
	if len(newH) > historyBufferSize {
		newH = newH[len(newH)-historyBufferSize:]
	}
	*h = newH
}

func (h History) AverageRatePerMS() float64 {
	var totalRate float64
	var usefulRates int
	for i, record := range h {
		if i == 0 {
			continue
		}
		previous := h[i-1]
		duration := record.At.Sub(previous.At)
		if duration.Milliseconds() == 0 {
			continue
		}
		rate := float64(record.Progress-previous.Progress) / float64(duration.Milliseconds())
		totalRate += rate
		usefulRates++
	}
	if usefulRates == 0 {
		return 0
	}
	return totalRate / float64(usefulRates)
}

type Record struct {
	Progress int64
	At       time.Time
}
