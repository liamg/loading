package bar

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/liamg/loading/pkg/util"
)

// Bar is a loading bar
type Bar struct {
	sync.Mutex
	w            *util.LineWriter
	renderFunc   RenderFunc
	total        int64
	current      int64
	linesMoved   int // number of lines moved down by the bar (0 if bar is drawn on same line)
	statsPadding int // number of spaces to pad the stats from the graphical bar component
	padding      int // number of spaces to the left and right of the bar
	label        string
	statsFuncs   []StatsFunc
	history      History
	complete     bool
	hideOnFinish bool
	logger       *bytes.Buffer
	autoComplete bool
	start        time.Time
}

// Default creates a loading bar with default settings
func Default() *Bar {
	return &Bar{
		start:      time.Now(),
		w:          util.NewLineWriter(os.Stderr),
		renderFunc: RenderSimple,
		statsFuncs: []StatsFunc{
			StatsPercentComplete,
			StatsAmountComplete,
			StatsAmountRate,
			StatsTimeRemaining,
		},
		statsPadding: 1,
		logger:       bytes.NewBuffer(nil),
		autoComplete: true,
	}
}

// New creates a bar with an optional list of option for customisation
func New(options ...Option) *Bar {
	b := Default()
	for _, option := range options {
		option(b)
	}
	return b
}

func (b *Bar) Log(format string, args ...interface{}) *Bar {
	b.Lock()
	defer b.Unlock()
	_, _ = b.logger.WriteString(fmt.Sprintf(format+"\n", args...))
	b.render()
	return b
}

// SetTotal sets the total value of the bar
func (b *Bar) SetTotal(total int) *Bar {
	b.SetTotalInt64(int64(total))
	return b
}

// SetLabel sets the label displayed at the start of the bar (this can be used to update the label during rendering)
func (b *Bar) SetLabel(label string) *Bar {
	b.Lock()
	defer b.Unlock()
	b.label = label
	b.render()
	return b
}

// SetTotalInt64 sets the total value of the bar with an int64
func (b *Bar) SetTotalInt64(total int64) *Bar {
	t := time.Now()
	b.Lock()
	defer b.Unlock()
	b.total = total
	b.updateHistory(t)
	b.render()
	return b
}

// AddCurrent adds the given value to the current value of the bar
func (b *Bar) AddCurrent(add int) *Bar {
	b.AddCurrentInt64(int64(add))
	return b
}

// Increment increases the current value by 1
func (b *Bar) Increment() {
	b.AddCurrentInt64(1)
}

// AddCurrentInt64 adds the given int64 value to the current value of the bar
func (b *Bar) AddCurrentInt64(current int64) *Bar {
	b.Lock()
	existing := b.current
	b.Unlock()
	b.SetCurrentInt64(existing + current)
	return b
}

// SetCurrent sets the current value of the bar
func (b *Bar) SetCurrent(current int) *Bar {
	b.SetCurrentInt64(int64(current))
	return b
}

// SetCurrentInt64 sets the current value of the bar with an int64
func (b *Bar) SetCurrentInt64(current int64) *Bar {
	t := time.Now()
	b.Lock()
	defer b.Unlock()
	b.current = current
	b.updateHistory(t)
	b.render()
	return b
}

func (b *Bar) updateHistory(t time.Time) {
	b.history.Push(Record{
		Progress: b.current,
		Duration: time.Since(b.start),
	})
}

// Write implements io.Writer on Bar, so it can be used to keep track of i/o operations using an io.MultiWriter
func (b *Bar) Write(p []byte) (n int, err error) {
	b.AddCurrent(len(p))
	return len(p), nil
}

// clean removes the bar from the terminal and places the cursor where rendering started
func (b *Bar) clean(w io.Writer) {

	// turn off cursor
	_, _ = w.Write([]byte("\x1b[?25l"))
	for line := 0; line < b.linesMoved; line++ {
		// move to start of line, clear line, move up one line
		_, _ = w.Write([]byte("\r\x1b[K\x1b[A"))
	}
	// move to start of line and clear it
	_, _ = w.Write([]byte("\r\x1b[K"))
}

func (b *Bar) getStats(completion float64) (before string, after string) {

	var segments []string
	for _, f := range b.statsFuncs {
		segments = append(segments, f(b.current, b.total, completion, b.history))
	}
	after = strings.Join(segments, " ")

	before = b.label
	if before != "" {
		before = before + strings.Repeat(" ", b.statsPadding)
	}
	if after != "" {
		after = strings.Repeat(" ", b.statsPadding) + after
	}
	return before, after
}

// Finish finishes processing on the bar, and restores the terminal state e.g. cursor visibility
// You don't need to call this unless:
//
//	A. You want to stop the bar before it completes
//	B. Your bar has an unknown (zero) total and thus cannot know when it is complete
func (b *Bar) Finish() {
	b.Lock()
	defer b.Unlock()
	b.finish()
}

// turn the cursor back on when finished
func (b *Bar) finish() {
	if !b.complete {
		if b.hideOnFinish {
			b.clean(b.w)
		}
		_, _ = b.w.Write([]byte("\x1b[?25h"))
		b.complete = true
	}
}

// render draw the bar to the terminal
func (b *Bar) render() {

	if b.complete {
		return
	}

	// find terminal width
	termWidth := util.TermWidth()

	var completion float64
	if b.total > 0 {
		// calculate the fraction of the bar that is filled
		completion = float64(b.current) / float64(b.total)
		if completion > 1 {
			completion = 1
		}
	}

	// write some stats
	statsBefore, statsAfter := b.getStats(completion)

	// calculate how much room we have left for the loading bar itself
	barWidth := termWidth - len(statsBefore) - len(statsAfter) - (b.padding * 2)
	if barWidth < 0 {
		barWidth = 0
	}

	// get the width to draw as filled
	completeWidth := int(math.Round(completion * float64(barWidth)))

	buf := bytes.NewBuffer(nil)

	// clean up the old render
	b.clean(buf)

	// reset the line counter - all rendering must be done below here
	b.w.Reset()

	_, _ = buf.WriteString(b.logger.String())
	_, _ = buf.WriteString("\x1b[0m" + strings.Repeat(" ", b.padding) + statsBefore)
	b.renderFunc(buf, completeWidth, barWidth)
	_, _ = buf.WriteString("\x1b[0m" + statsAfter + strings.Repeat(" ", b.padding))

	buf.WriteString("\n")

	_, _ = b.w.Write(buf.Bytes())
	b.linesMoved = b.w.Lines()

	// if we've finished, new line and turn on cursor
	if b.autoComplete && b.current == b.total && b.total > 0 {
		b.finish()
	}
}
