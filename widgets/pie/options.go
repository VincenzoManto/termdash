package pie 

import (
	"github.com/mum4k/termdash/cell"
)

type Option interface {
	// set sets the provided option.
	set(*options)
}

// option implements Option.
type option func(*options)

// set implements Option.set.
func (o option) set(opts *options) {
	o(opts)
}

// options stores the provided options.
type options struct{
	colors []cell.Color
}

// validates the provided options
// at the moment no validation is performed cause options are not required
func (o *options) validate() error {
	return nil
}

// newOptions creates a new options instance.
func newOptions() *options {
	return &options{
		colors: DefaultColors,
	}
}

var DefaultColors = []cell.Color{
	cell.ColorRed,
	cell.ColorGreen,
	cell.ColorBlue,
	cell.ColorYellow,
	cell.ColorMagenta,
	cell.ColorCyan,
	cell.ColorWhite,
}