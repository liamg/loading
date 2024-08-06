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
	latest := h[len(h)-1]

	remaining := total - latest.Progress
	if remaining <= 0 {
		return "00m00s"
	}

	duration := time.Duration(int64(float64(remaining)/avgRate)) * time.Millisecond
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
	if len(h) == 0 {
		return 0
	}
	latest := h[len(h)-1]
	return float64(latest.Progress) / float64(latest.Duration.Milliseconds())
}

type Record struct {
	Progress int64
	Duration time.Duration
}
