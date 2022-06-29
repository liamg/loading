package bar

// Option is a customisation function that modifies a Bar
type Option func(b *Bar)

// OptionWithRenderFunc sets the render function for the bar
func OptionWithRenderFunc(r RenderFunc) Option {
	return func(b *Bar) {
		b.renderFunc = r
	}
}

// OptionWithStatsPadding sets the number of spaces between the graphical bar and the stats segments
func OptionWithStatsPadding(padding int) Option {
	return func(b *Bar) {
		b.statsPadding = padding
	}
}

// OptionWithPadding sets the number of spaces to the left and right of the entire bar
func OptionWithPadding(padding int) Option {
	return func(b *Bar) {
		b.padding = padding
	}
}

// OptionWithoutStatsFuncs removes all stats functions from the bar
func OptionWithoutStatsFuncs() Option {
	return func(b *Bar) {
		b.statsFuncs = nil
	}
}

// OptionWithStatsFuncs sets the stats functions for the bar
func OptionWithStatsFuncs(f ...StatsFunc) Option {
	return func(b *Bar) {
		b.statsFuncs = f
	}
}

// OptionWithLabel sets the label for the bar (displayed to the left)
func OptionWithLabel(l string) Option {
	return func(b *Bar) {
		b.label = l
	}
}

// OptionHideOnFinish hides the bar when it is finished
func OptionHideOnFinish(enabled bool) Option {
	return func(b *Bar) {
		b.hideOnFinish = enabled
	}
}
