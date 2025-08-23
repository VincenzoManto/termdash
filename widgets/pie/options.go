package pie 

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
func (o *options) validate() error {
	if len(o.values) == 0 {
		return errors.New("values cannot be empty")
	}
	if len(o.colors) == 0 {
		return errors.New("colors cannot be empty")
	}
	for _, v := range o.values {
		if v < 0 {
			return errors.New("all values must be non-negative")
		}
	}
	return nil
}

// newOptions creates a new options instance.
func newOptions() *options {
	return &options{
		colors: DefaultColors,
	}
}

const DefaultColors = []cell.Color{
	cell.ColorRed,
	cell.ColorGreen,
	cell.ColorBlue,
	cell.ColorYellow,
	cell.ColorMagenta,
	cell.ColorCyan,
	cell.ColorWhite,
}